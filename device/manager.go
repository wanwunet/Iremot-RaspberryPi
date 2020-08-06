package device

import (
	"Iremot-RaspberryPi/models"
)

var GPIOMap map[int8]models.SetGpio = make(map[int8]models.SetGpio)

func Register(task Task) {
	task.Init()
	task.Run()
}
