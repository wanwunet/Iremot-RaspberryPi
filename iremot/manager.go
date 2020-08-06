package iremot

import (
	"Iremot-RaspberryPi/device"
	"Iremot-RaspberryPi/models"
	"time"
)

var MqttSingleton *MqttClientManger

func Open() {

	//MQTT连接
	MqttSingleton = NewMqttClient()
	MqttSingleton.NewConnect()
	MqttSingleton.Run()

	router()
	go Heartbeat()
}

func Close() {

}

//心跳
var HeartbeatID uint64

func Heartbeat() {
	var SleepTime int64 = 1000000000
	for {
		NowTime := time.Now().UnixNano()

		model := models.HeartbeatModel{
			Id:       HeartbeatID,
			SednTime: time.Now().Unix(),
			PinArr:   device.GetPinStateAll(),
		}

		MqttSingleton.MsgSend <- Message{TOPIC_HEARTBEAT(), model}
		//fmt.Println(device.GetPinStateAllJson())

		if diff := time.Now().UnixNano() - NowTime; diff < SleepTime {
			time.Sleep(time.Duration(SleepTime - diff))
		}

		HeartbeatID++
	}
}
