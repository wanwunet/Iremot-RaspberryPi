package iremot

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Pub Client Topic : %s \n", msg.Topic())
	fmt.Printf("Pub Client msg : %s \n", msg.Payload())
}

//创建全局mqtt sub消息处理 handler
var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Sub Client Topic : %s--", msg.Topic())
	fmt.Printf("Sub Client msg : %s \n", msg.Payload())
}

type message struct {
	topic   string
	payload interface{}
}

// 客户端管理器
type MqttClientManger struct {
	client   mqtt.Client
	msgSend  chan message
	topicSub []string
}

func newMqttClient(options *mqtt.ClientOptions) mqtt.Client {
	client := mqtt.NewClient(options)
	return client
}

func NewMqttClient() *MqttClientManger {

	clinetOptions := mqtt.NewClientOptions().AddBroker(Server)

	//设置用户名密码
	clinetOptions.SetUsername(Username).SetPassword(Password)
	//设置客户端ID
	clinetOptions.SetClientID(ClientID)
	//设置handler
	clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	//设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	//设置自动重连
	clinetOptions.SetAutoReconnect(true)
	//创建客户端连接
	c := newMqttClient(clinetOptions)
	msg := make(chan message)
	return &MqttClientManger{client: c, msgSend: msg}
}

// 客户端连接
func (mg *MqttClientManger) NewConnect() {
	if token := mg.client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fmt.Printf("[Pub] mqtt connect error, error: %s \n", token.Error())
		return
	}
}

// 发送消息
func (mg *MqttClientManger) Publish() {
	for {
		msg, ok := <-mg.msgSend
		if ok {
			// 格式化数据，将信息转换为json
			payload, err := json.Marshal(msg.payload)
			if err != nil {
				fmt.Println(err)
			}
			token := mg.client.Publish(msg.topic, 1, false, payload)
			token.Wait()
		}
	}

}

// 订阅消息
func (mg *MqttClientManger) Subscribe() {
	for _, topic := range mg.topicSub {
		token := mg.client.Subscribe(topic, 1, messageSubHandler)
		token.Wait()
	}
}

// 启动服务
func (mg *MqttClientManger) Run() {
	go mg.Subscribe()
	go mg.Publish()
}
