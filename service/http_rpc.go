//当前业务是否用到，未用到建议删除

package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	zlog "github.com/rs/zerolog/log"
)

type RequestBody struct {
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
	Secret  string            `json:"secret"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, body []byte, contentType string) (Response, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(body))
	var response Response
	if err != nil {
		zlog.Fatal().Msgf("请求异常 url(%s) 查找异常:%s", url, err.Error())
	} else {
		defer resp.Body.Close()
		result, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(result, &response)
		if err != nil {
			zlog.Fatal().Msgf("json数据解析异常(%s):%s", result, err.Error())
		}
		if response.Code != 0 {
			zlog.Fatal().Msgf("请求Code不为0(%s):%s", result, err.Error())
		}
	}
	return response, err
}
