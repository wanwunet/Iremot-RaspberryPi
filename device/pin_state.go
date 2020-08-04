package device

import (
	"Iremot-RaspberryPi/device/rpio"
	"encoding/json"
)

type PinState struct {
	Pin   int8
	Mod   uint32
	State uint32
}

func GetPinState(pin int8) (sta PinState) {
	sta.Pin = pin
	p := rpio.Pin(Physical2BCM[sta.Pin])
	sta.Mod = rpio.ReadPinMode(p)
	sta.State = uint32(rpio.ReadPin(p))
	return sta
}

func GetPinStateAll() (staList []PinState) {
	staList = make([]PinState, len(GpioPhysical))
	for i := 0; i < len(staList); i++ {
		staList[i] = GetPinState(GpioPhysical[i])
	}
	return staList
}

func GetPinStateAllJson() (str string) {
	staList := GetPinStateAll()
	if b, err := json.Marshal(staList); err == nil {
		str = string(b)
	}
	return str
}
