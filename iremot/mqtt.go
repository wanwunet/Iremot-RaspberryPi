package iremot

import (
	"Iremot-RaspberryPi/conf"
	"Iremot-RaspberryPi/util/md5"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//创建全局mqtt publish消息处理 handler
//var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("Pub Client Topic : %s \n", msg.Topic())
//	fmt.Printf("Pub Client msg : %s \n", msg.Payload())
//}

//创建全局mqtt sub消息处理 handler
//var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("Sub Client Topic : %s--", msg.Topic())
//	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
//}

type Message struct {
	Topic   string
	Payload interface{}
}

// 客户端管理器
type MqttClientManger struct {
	Client  mqtt.Client
	MsgSend chan Message
}

func NewMqttClient() *MqttClientManger {

	clinetOptions := mqtt.NewClientOptions().AddBroker(Server)

	//设置用户名密码
	pass := conf.Serial() + "-" + conf.Password
	clinetOptions.SetUsername(conf.ProductId).SetPassword(md5.MD5(pass))
	//设置客户端ID
	clinetOptions.SetClientID(conf.Ether())
	//设置handler
	//clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//设置自动重连
	clinetOptions.SetAutoReconnect(true)
	//创建客户端连接
	c := mqtt.NewClient(clinetOptions)
	msg := make(chan Message)
	return &MqttClientManger{Client: c, MsgSend: msg}
}

// 客户端连接
func (mg *MqttClientManger) NewConnect() {
	if token := mg.Client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Printf("[Pub] mqtt connect error, error: %s \n", token.Error())
		return
	}
}

// 发送消息
func (mg *MqttClientManger) Publish() {
	for {
		msg, ok := <-mg.MsgSend
		if ok {
			// 格式化数据，将信息转换为json
			payload, err := json.Marshal(msg.Payload)
			if err != nil {
				fmt.Println(err)
			}
			token := mg.Client.Publish(msg.Topic, 1, false, payload)
			token.Wait()
		}
	}

}

// 订阅消息
//func (mg *MqttClientManger) Subscribe() {
//	for _, topic := range mg.topicSub {
//		token := mg.client.Subscribe(topic, 1, messageSubHandler)
//		token.Wait()
//	}
//}

func (mg *MqttClientManger) Subscribe(topic string, messageSubHandler func(client mqtt.Client, msg mqtt.Message)) {
	mg.Client.Subscribe(topic, 1, messageSubHandler)
}

// 启动服务
func (mg *MqttClientManger) Run() {
	//go mg.Subscribe()
	go mg.Publish()
}
