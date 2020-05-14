package samples

import (
	"encoding/json"
	mc "tbClient/mqtt"
)

// 通过网关获取设备的客户端或者共享属性信息
func GateWayReqAtt(){
	done := make(chan bool, 1)
	//第二个参数是网关设备的token
	mc.GatewayMqttClient("tcp://120.25.167.36:1883","Q4JfMEAojK0Q2Vrm9Ke9","DXE500-000000001")
	att := map[string]interface{}{
		"id": 1,
		"device":"温度计1",
		"client":true,
		"key":"testyhb",
	}
	payloadAtt, _ := json.Marshal(att)
	subTopics := []string{
		"v1/gateway/attributes/response",
	}
	mc.GatewaySubscribeTopics(subTopics,gwSubRecvMessage)
	mc.GatewayPublishTopic("v1/gateway/attributes/request",1,payloadAtt)
	<-done
}
