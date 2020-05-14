package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	hq "tbClient/httpreq"
	"tbClient/util"
)

// iot相关http服务
// 基于thingsboard提供的协议

// 获取客户端/共享属性
// token attributetype keyname
// http://127.0.0.1:9090/iot/getDeviceAtt/7rSbzTHL6OlBWdjcexMz/1/testyhb
// 1:clientKey 客户端属性
// 2:sharedKeys 共享属性
func GetDeviceAtt(c *gin.Context){
	// 中间件或处理程序中启动新的Goroutines时（简单理解，另起线程），则不应该使用其中的原始上下文，你必须使用只读副本（c.Copy()）
	// cCp := c.Copy()
	tk := c.Param("tk")
	at := c.Param("at")
	kn := c.Param("kn")
	var iotUrl = ""
	if at == "1"{
		iotUrl = util.IotServer + "/api/v1/" + tk + "/attributes?clientKeys=" + kn
	}else{
		iotUrl = util.IotServer + "/api/v1/" + tk + "/attributes?sharedKeys=" + kn
	}
	//fmt.Println(iotUrl)
	// http://120.25.167.36:8080/api/v1/$ACCESS_TOKEN/attributes?clientKeys=attribute1,attribute2&sharedKeys=shared1,shared2
	resp := hq.HttpGet(iotUrl)
	c.String(http.StatusOK, "%s", resp)

}

// 增加/修改客户端属性
// http://localhost:9090/iot/postDeviceAtt/7rSbzTHL6OlBWdjcexMz/testyhb/7777777777
func PostDeviceAtt(c *gin.Context){

	tk := c.Param("tk")
	kn := c.Param("kn")
	val := c.Param("val")
	att := map[string]string{
		kn: val,
	}
	iotUrl := util.IotServer + "/api/v1/" + tk + "/attributes"
	payloadAtt, _ := json.Marshal(att)
	resp := hq.HttpPost(iotUrl, payloadAtt, "text/plain")
	c.String(http.StatusOK, "%s", resp)

}

type TelBody struct {
	// 首字母必须大写，否则无法绑定
	Keyname string `json:"keyname" binding:"required"`
	Keyvalue string `json:"keyvalue" binding:"required"`
}

// 上传单个遥测信息
// http://localhost:9090/iot/postDeviceTel/7rSbzTHL6OlBWdjcexMz
// Content-Type:application/json
// body:{"keyname":"test", "keyvalue":"45%"}
func PostDeviceTelemetry(c *gin.Context)  {

	tk := c.Param("tk")
	var reqInfo TelBody
	err := c.BindJSON(&reqInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Data Error"})
		return
	} else {
		tel := map[string]string{
			reqInfo.Keyname: reqInfo.Keyvalue,
		}
		iotUrl := util.IotServer + "/api/v1/" + tk + "/telemetry"
		payloadTel, _ := json.Marshal(tel)
		resp := hq.HttpPost(iotUrl, payloadTel, "text/plain")
		c.String(http.StatusOK, "%s", resp)
	}

	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//fmt.Println("fffffffffffff=========="+string(buf[0:n]))
}

// 批量上传遥测信息
// http://localhost:9090/iot/postDeviceMuiltTel/7rSbzTHL6OlBWdjcexMz
// Content-Type:application/json
// body:[{"keyname":"test", "keyvalue":"45%"},{"keyname":"fffff", "keyvalue":"65%"}......]
func PostDeviceMuiltTelemetry(c *gin.Context)  {

	tk := c.Param("tk")
	var telArr []TelBody   // Use slice for JSON array
	rd := json.NewDecoder(c.Request.Body)
	err := rd.Decode(&telArr)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Data Error"})
		return
	}else{
		for i:= 0;i<len(telArr);i++{
			tel := map[string]string{
				telArr[i].Keyname: telArr[i].Keyvalue,
			}
			iotUrl := util.IotServer + "/api/v1/" + tk + "/telemetry"
			payloadTel, _ := json.Marshal(tel)
			resp := hq.HttpPost(iotUrl, payloadTel, "text/plain")
			fmt.Println(resp)
		}
	}

}