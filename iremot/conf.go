package iremot

import "Iremot-RaspberryPi/conf"

//mqtt connect
var Server string = "tcp://192.168.1.80:1883"

//心跳
func TOPIC_HEARTBEAT() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.Ether() + "/heartbeat"
}

//设置gpio
func TOPIC_SETGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.Ether() + "/set_gpio"
}
