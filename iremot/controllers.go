package iremot

import (
	"Iremot-RaspberryPi/device"
	"Iremot-RaspberryPi/device/rpio"
	"Iremot-RaspberryPi/models"
	"Iremot-RaspberryPi/util/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//设置pin iremot/产品ID/设备ID/setgpio
var setgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.SetGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(device.Physical2BCM[mod.Pin])
	pin.Output()
	pin.Write(rpio.State(mod.State))
	fmt.Println(mod)
}

//设置pin iremot/产品ID/设备ID/getgpio
var getgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.GetGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(device.Physical2BCM[mod.Pin])
	pin.Input()
	fmt.Println(mod)
}

//设置pin iremot/产品ID/设备ID/getgpio
var pwmgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.PwmGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(device.Physical2BCM[mod.Pin])
	pin.Mode(rpio.Pwm)
	pin.Freq(mod.Freq)
	pin.DutyCycle(mod.DutyLen, mod.CycleLen)
	fmt.Println(mod)
}
