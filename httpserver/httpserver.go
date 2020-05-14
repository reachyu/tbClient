package httpserver

import (
	"log"
	"net/http"
	"tbClient/config"
	"strings"
	"github.com/gin-gonic/gin"
)

/*
开启http server
封装对外提供的接口服务
*/

var router *gin.Engine

// 不允许设备以外访问来源访问，只允许设备内应用访问
// 设备内访问，只能是http://127.0.0.1开头，localhost也不行
func checkHost() gin.HandlerFunc{
	return func(c *gin.Context){
		reqRemoteAddr := c.Request.RemoteAddr
		ipindex := strings.Contains(reqRemoteAddr,"127.0.0.1")
		if ipindex {
			c.Next()
		}else{
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			return
		}
	}
}

func init() {
	// Engin
	router = gin.Default()
	log.Println(">>>> IoT Server start <<<<")
	//router := gin.New()

	// 静态资源 http://127.0.0.1:9090/html/test.html
	router.Static("/html", "./assets")
}

func InitHttpServer() {
	// 全局中间件
	// 检测访问IP，判断是不是内部访问
	router.Use(checkHost())
	postListen()
	getListen()
	// 指定地址和端口号
	//router.Run("localhost:9090")
	router.Run(":9090")
}

//启动post服务监听
func postListen() {
	log.Println(">>>> start post servers <<<<")
	//post接口
	for u, f := range config.PostMaps {
		router.POST("/"+u, f)
	}
}

//启动get服务监听
func getListen() {
	log.Println(">>>> start get servers <<<<")
	//get接口
	for u, f := range config.GetMaps {
		router.GET("/"+u, f)
	}
}
