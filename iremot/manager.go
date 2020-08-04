package iremot

import (
	"Iremot-RaspberryPi/device"
	"time"
)

var MqttSingleton *MqttClientManger

func Open() {

	//MQTT连接
	MqttSingleton = NewMqttClient()
	MqttSingleton.NewConnect()
	MqttSingleton.Run()

	go Heartbeat()
}

func Close() {

}

//心跳
func Heartbeat() {

	var SleepTime int64 = 1000000000
	for {
		NowTime := time.Now().UnixNano()
		MqttSingleton.msgSend <- message{TOPIC_HEARTBEAT, device.GetPinStateAllJson()}
		//fmt.Println(drive.GetPinStateAllJson())

		if diff := time.Now().UnixNano() - NowTime; diff < SleepTime {
			time.Sleep(time.Duration(SleepTime - diff))
		}
	}
}
