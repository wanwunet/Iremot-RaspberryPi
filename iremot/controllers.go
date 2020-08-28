package iremot

import (
	"Iremot-RaspberryPi/device/pinmap"
	"Iremot-RaspberryPi/device/rpio"
	"Iremot-RaspberryPi/models"
	"Iremot-RaspberryPi/task"
	"Iremot-RaspberryPi/util/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//设置pin iremot/产品ID/设备ID/setgpio
var setgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.SetGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(pinmap.Physical2BCM[mod.Pin])
	pin.Output()
	pin.Write(rpio.State(mod.State))
	//fmt.Println(mod)
}

//设置pin iremot/产品ID/设备ID/getgpio
var getgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.GetGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(pinmap.Physical2BCM[mod.Pin])
	pin.Input()
	//fmt.Println(mod)
}

//设置pin iremot/产品ID/设备ID/pwmgpio
var pwmgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.PwmGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(pinmap.Physical2BCM[mod.Pin])
	pin.Mode(rpio.Pwm)
	pin.Freq(mod.Freq)
	pin.DutyCycle(mod.DutyLen, mod.CycleLen)
	//fmt.Println(mod)
}

//pin自定义序列 iremot/产品ID/设备ID/dataframegpio
var dataframegpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.DataFrameGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(pinmap.Physical2BCM[mod.Pin])
	pin.Output()
	for i := 0; i < len(mod.State); i++ {
		pin.Write(rpio.State(mod.State[i]))
		time.Sleep(time.Millisecond * time.Duration(mod.Delayed))
	}
	fmt.Println(mod)
}

//def gpio iremot/产品ID/设备ID/dataframegpio
var defgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var mod models.DefGpio
	json.Decode(msg.Payload(), &mod)
	pin := rpio.Pin(pinmap.Physical2BCM[mod.Pin])
	pin.Output()
	pin.Low()
	pin.Input()
	fmt.Println(mod)
}

//def gpio iremot/产品ID/设备ID/dataframegpio
var deviceallgpiotaskSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	mod := models.DeviceTask{}
	json.Decode(msg.Payload(), &mod)
	task.RegisterTask(mod)
}
