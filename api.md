# MQTT
在需要使用的go代码里import，作为sdk使用
## 获取设备属性
参见samples目录device_request_attributes.go、gateway_request_attributes.go

## 上传设备属性
参见samples目录device_publish_info.go、gateway_publish_info.go

## 上传设备遥测信息
参见samples目录device_subscribe_attribute.go、gateway_subscribe_attribute.go

# HTTP
* 已经限定只能127.0.0.1访问，如果不需要，请注释httpserver\httpserver.go中方法InitHttpServer()的代码router.Use(checkHost())
## 获取设备属性
### api url
http://ip:9090/iot/getDeviceAtt/token/attributetype/keyname
### 请求方式
get请求
### 参数说明
* token：thingsboard上设备token
* attributetype：1:clientKey 客户端属性   2:sharedKeys 共享属性
* keyname：属性名

## 上传设备属性
### api url
http://ip:9090/iot/postDeviceAtt/token/keyname/keyvalue
### 请求方式
post请求
### 参数说明
* token：thingsboard上设备token
* keyname：属性名
* keyvalue：属性值

## 上传单个遥测信息
### api url
http://ip:9090/iot/postDeviceTel/token
### 请求方式
post请求
### 参数说明
* token：thingsboard上设备token
* Content-Type=application/json
* body里传json串，{"keyname":"cpu", "keyvalue":"111111"}

## 批量上传遥测信息
### api url
http://ip:9090/iot/postDeviceMuiltTel/token
### 请求方式
post请求
### 参数说明
* token：thingsboard上设备token
* Content-Type=application/json
* body里传json串，[{"keyname":"cpu", "keyvalue":"45"},{"keyname":"mem", "keyvalue":"65"}]