package samples

import (
	"fmt"
	"encoding/json"
	mc "tbClient/mqtt"
)

// gateway订阅

// 订阅回调函数，处理接收到的订阅信息
func gwSubRecvMessage(msg mc.Message) {
	fmt.Printf("gateway Subscribe Topics recv msg %s\n", string(msg.Payload()))
}

// thingsboard官方指导，只有共享属性可以订阅
// delete:device Subscribe Topics recv msg {"deleted":["testgx"]}
// add:device Subscribe Topics recv msg {"kkkk":"43532"}
// update:device Subscribe Topics recv msg {"ggg":"466336"}
func GateWaySubTopics()  {
	done := make(chan bool, 1)
	//第二个参数是网关设备的token
	mc.GatewayMqttClient("tcp://120.25.167.36:1883","Q4JfMEAojK0Q2Vrm9Ke9","DXE500-000000001")
	subTopics := []string{
		"v1/gateway/attributes",
	}
	kv := map[string]string{
		"device":    "温度计1",
	}
	payloadTel, _ := json.Marshal(kv)
	mc.GatewayPublishTopic("v1/gateway/connect",1,payloadTel)
	mc.GatewaySubscribeTopics(subTopics,gwSubRecvMessage)
	<-done
}