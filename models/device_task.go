package models

import (
	"Iremot-RaspberryPi/device/pinmap"
	"Iremot-RaspberryPi/device/rpio"
	"fmt"
	"time"
)

type DeviceTask struct {
	Id        uint64
	Name      string
	Loop      int64
	TaskFrame []DataFrame
}

type DataFrame struct {
	PinSta  []DataFramePinSta
	Delayed uint64
}

type DataFramePinSta struct {
	Pin int
	Sta int
}

func (this *DeviceTask) Run() {
	for i := int64(0); i < this.Loop; i++ {

		for j := 0; j < len(this.TaskFrame); j++ {
			frame := this.TaskFrame[j]
			if frame.Delayed > 0 {
				time.Sleep(time.Duration(frame.Delayed))
			}
			for k := 0; k < len(frame.PinSta); k++ {
				p := rpio.Pin(pinmap.Physical2BCM[int8(frame.PinSta[k].Pin)])
				p.Write(rpio.State(frame.PinSta[k].Sta))
				fmt.Println(i, j, k, frame.PinSta[k].Pin, frame.PinSta[k].Sta)
			}
		}

	}
}

func (this *DeviceTask) CallRun(call func(Id uint64)) {
	this.Run()
	call(this.Id)
}
