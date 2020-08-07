package iremot

func router() {
	MqttSingleton.Subscribe(TOPIC_SETGPIO(), setgpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_GETGPIO(), getgpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_PWMGPIO(), pwmgpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_DATAFRAMEGPIO(), dataframegpioSubHandler)
	MqttSingleton.Subscribe(TOPIC_DEFGPIO(), defgpioSubHandler)
}
