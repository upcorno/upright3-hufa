package service

import (
	"law/conf"
	"net/http"

	zlog "github.com/rs/zerolog/log"

	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

type wxSrv struct {
}

var WxSrv *wxSrv = &wxSrv{}

func (w *wxSrv) WxNotify(repWriter http.ResponseWriter, request *http.Request) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	handler := func(req map[string]interface{}) map[string]interface{} {
		//暂时没有需要处理的微信通知，先日志记录
		zlog.Info().Msgf("wechat notify: %v", req)
		return nil
	}
	srv, err := sdk.NewServer(conf.App.WxApp.NotifyToken, conf.App.WxApp.NotifyAesKey, conf.App.WxApp.NotifyMchId, conf.App.WxApp.NotifyApiKey, true, handler)
	if err != nil {
		zlog.Error().Msgf("init wecat notify server error: %s", err)
	}
	if err := srv.Serve(repWriter, request); err != nil {
		zlog.Error().Msgf("wecat notify server serving error: %s", err)
	}
}

func (w *wxSrv) wxLogin(code string) (*weapp.LoginResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	return sdk.Login(code)
}

func (w *wxSrv) getPhoneNumber(code string) (*phonenumber.GetPhoneNumberResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	cli := sdk.NewPhonenumber()
	return cli.GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: code})
}
