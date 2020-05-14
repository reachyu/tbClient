package samples

import (
	"encoding/json"
	"tbClient/httpreq"
	"fmt"
	)

func GetDeviceAtt()  {
	url := "http://127.0.0.1:9090/iot/getDeviceAtt/7rSbzTHL6OlBWdjcexMz/1/testyhb"
	resp := httpreq.HttpGet(url)
	fmt.Println(resp)
}

func PostDeviceAtt(){

	url := "http://127.0.0.1:9090/iot/postDeviceAtt/7rSbzTHL6OlBWdjcexMz/testyhb/7777777777"
	resp := httpreq.HttpPost(url,nil,"text/plain")
	fmt.Println(resp)
}

func PostDeviceTelemetry(){
	url := "http://127.0.0.1:9090/iot/postDeviceTel/7rSbzTHL6OlBWdjcexMz"
	tel := map[string]string{
		"keyname":"cpu",
		"keyvalue":"1234567890",
	}
	data, _ := json.Marshal(tel)
	resp := httpreq.HttpPost(url,data,"application/json")
	fmt.Println(resp)
}

func PostDeviceMuiltTelemetry(){
	url := "http://127.0.0.1:9090/iot/postDeviceMuiltTel/7rSbzTHL6OlBWdjcexMz"
	tel := map[string]string{
		"keyname":"cpu",
		"keyvalue":"11111111111",
	}
	data, _ := json.Marshal(tel)
	resp := httpreq.HttpPost(url,data,"application/json")
	fmt.Println(resp)
}