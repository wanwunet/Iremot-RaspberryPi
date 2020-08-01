package drive

import (
	"Iremot-RaspberryPi/drive/rpio"
)

type PinState struct {
	Pin   uint8
	Mod   uint8
	State uint8
}

func GetPinState(pin uint8) (sta PinState) {
	sta.Pin = pin
	p := rpio.Pin(sta.Pin)
	sta.Mod = uint8(rpio.ReadPinMode(p))
	sta.State = uint8(rpio.ReadPin(p))
	return sta
}

func GetPinStateAll() (staList []PinState) {
	staList = make([]PinState, len(GpioPhysical))
	for i := 0; i < len(staList); i++ {
		staList[i] = GetPinState(uint8(Physical2BCM[GpioPhysical[i]]))
	}
	return staList
}
