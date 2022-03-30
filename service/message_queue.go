//当前业务是否用到，未用到建议删除
package service

import (
	"encoding/json"
	"law/conf"
)

func Send(body RequestBody) {
	rpc := conf.App.Rpc
	body.Secret = rpc.Secret
	jsonStr, _ := json.Marshal(body)
	Post(rpc.MessageQueryPushUrl, jsonStr, "application/json")
}
