package device

import (
	"Iremot-RaspberryPi/device/rpio"
	"Iremot-RaspberryPi/models"
	"encoding/json"
)

func GetPinState(pin int8) (sta models.PinState) {
	sta.Pin = pin
	p := rpio.Pin(Physical2BCM[sta.Pin])
	sta.Mod = rpio.ReadPinMode(p)
	sta.State = uint32(rpio.ReadPin(p))
	return sta
}

func GetPinStateAll() (staList []models.PinState) {
	staList = make([]models.PinState, len(GpioPhysical))
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
