package service

import (
	"context"
	"fmt"
	"law/conf"
	dao "law/dao"
	"log"
	"net/http"
	"os"

	zlog "github.com/rs/zerolog/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type wxSrv struct {
}

var WxSrv *wxSrv = &wxSrv{}

var mini *miniprogram.MiniProgram
var official *officialaccount.OfficialAccount

func init() {
	wechatLogFile, err := os.OpenFile("logs/wechat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件logs/wechat.log")
	}
	log.SetOutput(wechatLogFile)
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        conf.App.Rdb.RdbHost + ":" + fmt.Sprint(conf.App.Rdb.RdbPort),
		Database:    conf.App.Rdb.DbIndex,
		Password:    conf.App.Rdb.RdbPasswd,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	cfg := &miniConfig.Config{
		AppID:     conf.App.WxApp.Appid,
		AppSecret: conf.App.WxApp.Secret,
		Cache:     redisCache,
	}
	mini = wc.GetMiniProgram(cfg)

	officialAccountCfg := &offConfig.Config{
		AppID:          conf.App.WxApp.Appid,
		AppSecret:      conf.App.WxApp.Secret,
		Token:          conf.App.WxApp.NotifyToken,
		EncodingAESKey: conf.App.WxApp.NotifyAesKey,
		Cache:          redisCache,
	}
	official = wc.GetOfficialAccount(officialAccountCfg)
}

func (w *wxSrv) WxNotify(request *http.Request, repWriter http.ResponseWriter) {
	server := official.GetServer(request, repWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		switch msg.Event {
		case message.EventSubscribeMsgPopupEvent:
			openid := string(msg.FromUserName)
			uid, err := UserSrv.getUidAndSync(openid, "")
			if err != nil {
				zlog.Error().Msgf("wechat notify server serving error: %s", err.Error())
				break
			}
			for _, ev := range msg.GetSubscribeMsgPopupEvents() {
				if ev.SubscribeStatusString == "accept" {
					templateId := ev.TemplateID
					err = dao.TMsgSubDao.IncrSubscribeNum(uid, templateId)
					if err != nil {
						zlog.Error().Msgf("wechat notify server serving error: %s", err.Error())
						break
					}
				}
			}
		}
		text := message.NewText("success")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	err := server.Serve()
	if err != nil {
		zlog.Error().Msgf("wechat notify server serving error: %s", err.Error())
		return
	}
	server.Send()
	if err != nil {
		zlog.Error().Msgf("wechat notify server serving error: %s", err.Error())
		return
	}
}

func (w *wxSrv) wxLogin(code string) (openid string, unionid string, err error) {
	res, err := mini.GetAuth().Code2Session(code)
	if err != nil {
		return
	}
	openid = res.OpenID
	unionid = res.UnionID
	return
}

func (w *wxSrv) getPhoneNumber(code string) (phoneNumber string, err error) {
	rep, err := mini.GetAuth().GetPhoneNumber(code)
	if err != nil {
		return
	}
	phoneNumber = rep.PhoneInfo.PurePhoneNumber
	return
}
