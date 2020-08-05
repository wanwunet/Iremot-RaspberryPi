package iremot

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//心跳 iremot/产品ID/设备ID/setgpio
var setgpioSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	top := msg.Topic()
	mes := string(msg.Payload())
	fmt.Println(top, "-----", mes)
}
