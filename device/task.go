package device

import (
	"Iremot-RaspberryPi/device/rpio"
	"fmt"
	"time"
)

type Task interface {
	Init()
	Run()
}

type OutputTask struct {
	Task
	Pin     int8
	State   uint32
	Loop    int
	Delayed int64
}

func (this *OutputTask) Init() {
	fmt.Println("init")
}

func (this *OutputTask) Run() {

	index := 0
	pin := rpio.Pin(Physical2BCM[this.Pin])

START:

	NowTime := time.Now().UnixNano() / 1e6

	if this.State == 1 {
		pin.Output()
		pin.High()
	} else if this.State == 2 {
		pin.Output()
		pin.Low()
	} else if this.State == 3 {
		pin.PullUp()
	} else if this.State == 4 {
		pin.PullDown()
	}

	if diff := time.Now().UnixNano()/1e6 - NowTime; diff < this.Delayed {
		time.Sleep(time.Duration(this.Delayed-diff) * time.Millisecond)
	}

	if this.Loop == -1 || index < this.Loop {
		index++
		goto START
	}
}
