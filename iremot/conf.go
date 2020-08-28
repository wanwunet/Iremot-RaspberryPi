package iremot

import "Iremot-RaspberryPi/conf"

//mqtt connect
var Server string = "tcp://192.168.1.80:1883"

//心跳
func TOPIC_HEARTBEAT() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/heartbeat"
}

//输出 gpio
func TOPIC_SETGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/set_gpio"
}

//输入 gpio
func TOPIC_GETGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/get_gpio"
}

//PWM gpio
func TOPIC_PWMGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/pwm_gpio"
}

//DATA FRAME gpio
func TOPIC_DATAFRAMEGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/datagrame_gpio"
}

//def gpio
func TOPIC_DEFGPIO() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.ID() + "/def_gpio"
}

//def gpio
func TOPIC_DEVICEALLGPIOTASK() string {
	return "iremot_pi/" + conf.ProductId + "/device_all/gpio_task"
}
