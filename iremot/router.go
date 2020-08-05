package iremot

func router() {

	MqttSingleton.Subscribe("iremot_pi/+/+/set_gpio", setgpioSubHandler)

}
