package samples

import (
	"encoding/json"
	mc "tbClient/mqtt"
)

// 从server请求属性

// 获取设备的客户端或者共享属性信息
func ReqAtt(){
	done := make(chan bool, 1)
	mc.DeviceMqttClient("tcp://120.25.167.36:1883","7rSbzTHL6OlBWdjcexMz","DXE500-000000001")
	att := map[string]string{
		"clientKeys": "testyhb",
		"sharedKeys":"testgx",
	}
	payloadAtt, _ := json.Marshal(att)
	subTopics := []string{
		"v1/devices/me/attributes/response/+",
	}
	mc.SubscribeTopics(subTopics,subRecvMessage)
	mc.PublishTopic("v1/devices/me/attributes/request/1",1,payloadAtt)
	<-done
}
