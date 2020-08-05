package iremot

import "Iremot-RaspberryPi/conf"

//mqtt connect
var Server string = "tcp://192.168.1.80:1883"

//topic
func TOPIC_HEARTBEAT() string {
	return "iremot_pi/" + conf.ProductId + "/" + conf.Ether() + "/heartbeat"
}
