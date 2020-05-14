package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
	"time"
)

// 基于thingsboard提供的协议

const (
	keepalive   = 10 //keepalive ttl
	connTimeout = 5
	msgChnSize  = 20
)
var lock sync.RWMutex
var client mqtt.Client
var msgChl = make(chan Message, msgChnSize)

// Message type exported from MQTT client
type Message mqtt.Message
// MessageCB is Callback function type for receiving mqtt message
type MessageCB func(Message)
type FuncConnectionLost func(mqtt.Client, error)

// tcp://192.168.1.1:1883
// 7rSbzTHL6OlBWdjcexMz  (device token)
// DXE500-000000001
func DeviceMqttClient(url string, username string, clientid string) {
	options := mqtt.NewClientOptions().SetUsername(username).SetAutoReconnect(false).SetClientID(clientid).AddBroker(url).SetKeepAlive(keepalive * time.Second).SetWriteTimeout(connTimeout * time.Second).SetConnectionLostHandler(
		func(c mqtt.Client, err error) {
			log.Println("Mqtt client lost connection, error:", err)
			var cb FuncConnectionLost
			if cb != nil {
				cb(c, err)
			}
		})

	lock.Lock()
	if client == nil {
		client = mqtt.NewClient(options)
	}
	lock.Unlock()

	if client == nil || client.IsConnected() {
		fmt.Println("connect mqtt return,", client.IsConnected())
	}

	if token := client.Connect(); token.WaitTimeout(3*time.Second) && token.Error() != nil {
		panic(token.Error())
	}

}

// Publish will publish a message with the specified QoS and content
// to the specified topic.
// Returns a error
func PublishTopic(topic string, qos byte, msg []byte) error{

	if client == nil || !client.IsConnected() {
		log.Println("(publish)Not connected")
		return nil
	}

	token := client.Publish(topic, qos, false, msg)
	if token.WaitTimeout(5*time.Second) == false {
		return token.Error()
	}

	fmt.Printf("PublishTopic parameter %s\n", string(msg))

	return nil
}

// Subscribe starts a new subscription. Provide a MessageHandler to be executed when
// a message is published on the topic provided, or nil for the default handler
func SubscribeTopics(topics []string,cb MessageCB) {
	log.Println("subscribe:", topics)

	qos := byte(2)
	for i := 0; i < len(topics); i++ {
		tk := client.Subscribe(topics[i], qos,
			func(c mqtt.Client, msg mqtt.Message) {
				cb(msg)
			})

		if tk.Wait() && tk.Error() != nil {
			fmt.Println(tk.Error())
			break
		}
	}
}

//disconnect mqtt server
func Disconnect() {
	if client != nil && client.IsConnected() {
		log.Println("Disconnect")
		client.Disconnect(10)
		lock.Lock()
		client = nil
		lock.Unlock()
	}
}