# 简介
* 基于[Eclipse Paho MQTT Go client](https://github.com/eclipse/paho.mqtt.golang)、GIN框架实现ThingsBoard提供的MQTT、HTTP API
* 设备功能：上传设备遥测信息、上传设备属性、订阅信息
* 网关功能：上传设备遥测信息、上传设备属性、订阅信息
* api使用说明请见api.md
# 目录介绍
## mqtt
ThingsBoard提供的MQTT API实现
## httpservice
ThingsBoard提供的HTTP API实现
## httpserver
GIN框架启动http服务实现
## httpreq
通用http请求实现
## config
配置文件
svcconfig.go配置httpservice目录提供的http服务，GIN框架启动的时候会监听这些配置的http服务
## samples
ThingsBoard的MQTT、HTTP API服务使用sample

# 编译
```
go build
```

# 启动
```
go run main.go
```