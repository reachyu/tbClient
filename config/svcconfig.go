package config

import (
	hs "tbClient/service"

	"github.com/gin-gonic/gin"
)

/*
配置对外提供服务的访问名及其对应的实现方法
*/
var GetMaps = map[string]gin.HandlerFunc{
	// token attributetype keyname
	"iot/getDeviceAtt/:tk/:at/:kn": hs.GetDeviceAtt,
}

var PostMaps = map[string]gin.HandlerFunc{
	// token keyname value
	"iot/postDeviceAtt/:tk/:kn/:val": hs.PostDeviceAtt,
	// request body
	"iot/postDeviceTel/:tk": hs.PostDeviceTelemetry,
	"iot/postDeviceMuiltTel/:tk": hs.PostDeviceMuiltTelemetry,
}
