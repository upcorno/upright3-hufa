package service

import (
	"context"
	"fmt"
	"law/conf"
	"law/dao"
	"net/http"
	"time"

	zlog "github.com/rs/zerolog/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
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

func (w *wxSrv) SendConsulNotify(userId int, consulId int, question string, consulCreateTime int) {
	pageUrl := "/pages/consultant/Consultant?consultation_id=" + fmt.Sprint(consulId)
	data := map[string]*subscribe.DataItem{
		"thing6": {Value: question},
		"date7":  {Value: time.Unix(int64(consulCreateTime), 0).Format("2006-01-02 15:04")},
		"date4":  {Value: time.Now().Format("2006-01-02 15:04")},
	}
	w.sendNotify(userId, conf.App.WxApp.TemplateIdConsul, pageUrl, data)
}

func (w *wxSrv) sendNotify(userId int, templateId string, page string, data map[string]*subscribe.DataItem) {
	result, err := dao.TMsgSubDao.DecrSubscribeNum(userId, templateId)
	if err != nil {
		return
	}
	has, user, err := dao.UserDao.Get(userId, "", "")
	if err != nil {
		return
	}
	if result && has {
		msg := &subscribe.Message{
			ToUser:     user.Openid,
			TemplateID: templateId,
			Data:       data,
			Page:       page,
		}
		err = mini.GetSubscribe().Send(msg)
		if err != nil {
			zlog.Error().Err(err).Msg("wechat send notify error.")
			return
		}
	}
}

func (w *wxSrv) WxNotify(request *http.Request, repWriter http.ResponseWriter) {
	server := official.GetServer(request, repWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		switch msg.Event {
		case message.EventSubscribeMsgPopupEvent:
			openid := string(msg.FromUserName)
			uid, err := UserSrv.sync(openid, "")
			if err != nil {
				zlog.Error().Err(err).Msg("wechat notify server serving error.")
				break
			}
			for _, ev := range msg.GetSubscribeMsgPopupEvents() {
				if ev.SubscribeStatusString == "accept" {
					templateId := ev.TemplateID
					err = dao.TMsgSubDao.IncrSubscribeNum(uid, templateId)
					if err != nil {
						zlog.Error().Err(err).Msg("wechat notify server serving error.")
						break
					}
				}
			}
		default:

		}
		text := message.NewText("success")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	err := server.Serve()
	if err != nil {
		zlog.Error().Err(err).Msg("wechat notify server serving error.")
		return
	}
	server.Send()
	if err != nil {
		zlog.Error().Err(err).Msg("wechat notify server serving error.")
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
