package service

import (
	"law/conf"
	"law/model"
	"log"
	"net/http"
	"os"

	zlog "github.com/rs/zerolog/log"

	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/logger"
	"github.com/medivhzhan/weapp/v3/phonenumber"
	"github.com/medivhzhan/weapp/v3/server"
)

type wxSrv struct {
}

var WxSrv *wxSrv = &wxSrv{}

var wxServer *server.Server
var wxClient *weapp.Client

func init() {
	wechatFile, err := os.OpenFile("logs/wechat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件logs/wechat.log")
	}
	lgr := logger.NewLogger(log.New(wechatFile, "\r\n", log.LstdFlags), logger.Info, true)
	wxClient = weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret, weapp.WithLogger(lgr))
	handler := func(req map[string]interface{}) map[string]interface{} {
		//暂时没有需要处理的微信通知，先日志记录
		zlog.Info().Msgf("wechat notify: %v", req)
		return nil
	}
	wxServer, err = wxClient.NewServer(conf.App.WxApp.NotifyToken, conf.App.WxApp.NotifyAesKey, conf.App.WxApp.NotifyMchId, conf.App.WxApp.NotifyApiKey, true, handler)
	if err != nil {
		zlog.Error().Msgf("init wecat notify server error: %s", err)
	}
	wxServer.OnSubscribeMsgPopup(func(popupEv *server.SubscribeMsgPopupEvent) {
		openid := popupEv.FromUserName
		uid, err := UserSrv.GetUidAndSync(openid, "")
		if err != nil {
			return
		}
		for _, v := range popupEv.SubscribeMsgPopupEvent {
			if v.SubscribeStatusString == "accept" {
				templateId := v.TemplateId
				model.TMsgSubDao.IncrSubscribeNum(uid, templateId)
			}
		}
	})
}

func (w *wxSrv) WxNotify(repWriter http.ResponseWriter, request *http.Request) {
	if err := wxServer.Serve(repWriter, request); err != nil {
		zlog.Error().Msgf("wecat notify server serving error: %s", err)
	}
}

func (w *wxSrv) wxLogin(code string) (*weapp.LoginResponse, error) {
	return wxClient.Login(code)
}

func (w *wxSrv) getPhoneNumber(code string) (*phonenumber.GetPhoneNumberResponse, error) {
	return wxClient.NewPhonenumber().GetPhoneNumber(
		&phonenumber.GetPhoneNumberRequest{Code: code})
}
