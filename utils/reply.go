package utils

import (
	zlog "github.com/rs/zerolog/log"
)

// Reply  format
type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// page format
// Message
type page struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}

///client error  1开头
///server error  2开头
///unknown error  3开头
const (
	stSucc    int = 0    //正常
	stErrIpt  int = 1001 //输入数据有误
	stErrDeny int = 1008 //没有权限
	stErrJwt  int = 1004 //jwt未通过验证
	stErrOpt  int = 2001 //无数据返回
	stErrSvr  int = 2002 //服务端错误
	stFail    int = 3001 //失败
	stExt     int = 4000 //其他约定
)

func newReply(code int, msg string, data ...interface{}) (int, Reply) {
	if code != stSucc {
		zlog.Info().Msgf("msg:%s\ndetail:%v", msg, data)
		zlog.Warn().Msgf("msg:%s\ndetail:%v", msg, data)
		//当返回值不正常时主动将日志刷新至文件,以便分析
		FlushLog()
	}
	if len(data) > 0 {
		return 200, Reply{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
	}
	return 200, Reply{
		Code: code,
		Msg:  msg,
	}
}

// Succ 返回一个成功标识的结果格式
func Succ(msg string, data ...interface{}) (int, Reply) {
	return newReply(stSucc, msg, data...)
}

// Fail 返回一个失败标识的结果格式
func Fail(msg string, data ...interface{}) (int, Reply) {
	return newReply(stFail, msg, data...)
}

// Page 返回一个带有分页数据的结果格式
func Page(msg string, items interface{}, count int) (int, Reply) {
	return 200, Reply{
		Code: stSucc,
		Msg:  msg,
		Data: page{
			Items: items,
			Count: count,
		},
	}
}

// ErrIpt 返回一个输入错误的结果格式
func ErrIpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrIpt, msg, data...)
}

// ErrOpt 返回一个输出错误的结果格式
func ErrOpt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrOpt, msg, data...)
}

// ErrDeny 返回一个没有权限的结果格式
func ErrDeny(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrDeny, msg, data...)
}

// ErrJwt 返回一个未通过验证的结果格式
func ErrJwt(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrJwt, msg, data...)
}

// ErrSvr 返回一个服务端错误的结果格式
func ErrSvr(msg string, data ...interface{}) (int, Reply) {
	return newReply(stErrSvr, msg, data...)
}

// Ext 返回一个其他约定的结果格式
func Ext(msg string, data ...interface{}) (int, Reply) {
	return newReply(stExt, msg, data...)
}
