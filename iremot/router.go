package iremot

func router() {

	MqttSingleton.Subscribe(TOPIC_SETGPIO(), setgpioSubHandler)

}
