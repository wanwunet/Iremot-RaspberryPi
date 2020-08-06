package iremot

func router() {

	MqttSingleton.Subscribe(TOPIC_SETGPIO(), setgpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_GETGPIO(), getgpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_PWMGPIO(), pwmgpioSubHandler)
}
