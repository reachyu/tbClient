package samples

import (
	"encoding/json"
	"time"
	mc "tbClient/mqtt"
)

// 上传设备属性、遥测信息

// 上传设备属性信息
func PublishAtt(){
	mc.DeviceMqttClient("tcp://120.25.167.36:1883","7rSbzTHL6OlBWdjcexMz","DXE500-000000001")
	att := map[string]string{
		"testyhb": "20201010",
	}
	payloadAtt, _ := json.Marshal(att)
	mc.PublishTopic("v1/devices/me/attributes",2,payloadAtt)
	mc.Disconnect()
}

// 上传设备遥测信息
func PublishTel()  {
	mc.DeviceMqttClient("tcp://120.25.167.36:1883","7rSbzTHL6OlBWdjcexMz","DXE500-000000001")
	tel := map[string]int{
		"cpu":    11111,
		"mem": 22222,
	}
	payloadTel, _ := json.Marshal(tel)
	mc.PublishTopic("v1/devices/me/telemetry",2,payloadTel)
	mc.Disconnect()
}

// 上传设备遥测信息带时间戳
func PublishTelWithTS()  {
	mc.DeviceMqttClient("tcp://120.25.167.36:1883","7rSbzTHL6OlBWdjcexMz","DXE500-000000001")
	tel := map[string]int{
		"cpu":    33,
		"mem": 44,
	}
	telts := map[string]interface{}{
		"ts":time.Now().Unix(),
		"values":tel,
	}
	payloadTel, _ := json.Marshal(telts)
	mc.PublishTopic("v1/devices/me/telemetry",2,payloadTel)
	mc.Disconnect()
}