package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 基于thingsboard提供的协议

var gwclient mqtt.Client

// tcp://192.168.1.1:1883
// 7rSbzTHL6OlBWdjcexMz  (device token)
// DXE500-000000001
func GatewayMqttClient(url string, username string, clientid string) {
	DeviceMqttClient(url, username, clientid)
}

// Publish will publish a message with the specified QoS and content
// to the specified topic.
// Returns a error
func GatewayPublishTopic(topic string, qos byte, msg []byte) error{
	return PublishTopic(topic, qos, msg)
}

// Subscribe starts a new subscription. Provide a MessageHandler to be executed when
// a message is published on the topic provided, or nil for the default handler
func GatewaySubscribeTopics(topics []string,cb MessageCB) {
	SubscribeTopics(topics,cb)
}

// disconnect mqtt server
func GatewayDisconnect() {
	Disconnect()
}