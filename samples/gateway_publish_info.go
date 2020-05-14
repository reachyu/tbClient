package samples

import (
	"encoding/json"
	"time"
	mc "tbClient/mqtt"
)

// 通过网关上传设备属性信息
func GateWayPublishAtt(){
	//第二个参数是网关设备的token
	mc.GatewayMqttClient("tcp://120.25.167.36:1883","Q4JfMEAojK0Q2Vrm9Ke9","DXE500-000000001")
	kv := map[string]string{
		"testyhb": "90909090",
	}
	att := map[string]interface{}{
		"温度计1":kv,
	}
	payloadAtt, _ := json.Marshal(att)
	mc.GatewayPublishTopic("v1/gateway/attributes",2,payloadAtt)
	mc.GatewayDisconnect()
}

// 通过网关上传设备遥测信息带时间戳
func GateWayPublishTelWithTS()  {
	//第二个参数是网关设备的token
	mc.GatewayMqttClient("tcp://120.25.167.36:1883","Q4JfMEAojK0Q2Vrm9Ke9","DXE500-000000001")
	tel := map[string]int{
		"cpu":    9999,
		"mem": 77777,
	}
	telts := map[string]interface{}{
		"ts":time.Now().Unix(),
		"values":tel,
	}
	teltslist := []interface{}{telts}
	kv := map[string]interface{}{
		"温度计1":    teltslist,
	}
	payloadTel, _ := json.Marshal(kv)
	mc.GatewayPublishTopic("v1/gateway/telemetry",2,payloadTel)
	mc.GatewayDisconnect()
}