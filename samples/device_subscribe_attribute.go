package samples

import (
	"fmt"
	mc "tbClient/mqtt"
)

// 设备订阅

// 订阅回调函数，处理接收到的订阅信息
func subRecvMessage(msg mc.Message) {
	fmt.Printf("device Subscribe Topics recv msg %s\n", string(msg.Payload()))
}

// thingsboard官方指导，只有共享属性可以订阅
// delete:device Subscribe Topics recv msg {"deleted":["testgx"]}
// add:device Subscribe Topics recv msg {"kkkk":"43532"}
// update:device Subscribe Topics recv msg {"ggg":"466336"}
func SubTopics()  {
	done := make(chan bool, 1)
	mc.DeviceMqttClient("tcp://120.25.167.36:1883","7rSbzTHL6OlBWdjcexMz","DXE500-000000001")
	subTopics := []string{
		"v1/devices/me/attributes",
	}
	mc.SubscribeTopics(subTopics,subRecvMessage)
	<-done
}