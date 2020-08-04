package iremot

//mqtt connect
var Server string = "tcp://192.168.1.80:1883"
var Username string = "100000"
var Password string = "123456"
var ClientID string = "b8:27:eb:ef:da:e0"

//topic
var TOPIC_HEARTBEAT = "iremot/" + Username + "/" + ClientID + "/heartbeat"
