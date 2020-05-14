package httpreq

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
提供http请求通用服务
*/

//发送GET请求
//url:请求地址
//response:请求返回的内容
func HttpGet(httpurl string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(httpurl)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	response = string(result)
	return
}

//发送POST请求提交表单，可以设置header
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求返回的内容
func HttpPostForm(httpurl string, postString string, contentType string) (content string) {
	/*
		postValue := url.Values{
			"email":    {"xx@xx.com"},
			"password": {"123456"},
		}
		postString := postValue.Encode()
	*/
	req, err := http.NewRequest("POST", httpurl, strings.NewReader(postString))
	req.Header.Add("Content-Type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// 表单方式(必须)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//AJAX 方式请求
	//req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)

	return
}

//发送POST请求，可以设置header
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求返回的内容
func HttpPost(httpurl string, data []byte, contentType string) (content string) {
	/*
		tel := map[string]string{
			"testyhb": "et43w",
		}
		data, _ := json.Marshal(tel)
	*/
	req, err := http.NewRequest("POST", httpurl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)

	return
}
