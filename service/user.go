package service

import (
	"errors"
	"law/conf"
	"law/model"
	"law/utils"
	"time"

	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

type user struct{}

var User *user

func (u *user) Login(code string, ip string) (token string, err error) {
	res, err := u.wxLogin(code)
	if err != nil {
		return
	}
	if err = res.GetResponseError(); err != nil {
		return
	}
	uid, err := u.getUid(res.OpenID, res.UnionID)
	if err != nil {
		return
	}
	return utils.CreateAuthToken(uid, ip)
}

// 根据 openId 获取用户 id，不存在时创建新用户返回对应 id
func (u *user) getUid(openid string, unionID string) (uid int, err error) {
	user := &model.User{Openid: openid}
	has, err := user.Get()
	if err != nil {
		return
	}
	if has {
		uid = user.Id
		if user.Unionid != unionID {
			//更新 unionid, unionid 可能从 null 变为由内容
			err = user.Update()
		}
		return
	}
	user.AppId = conf.App.WxApp.Appid
	user.Unionid = unionID
	user.CreateTime = int(time.Now().Unix())
	if err = user.Insert(); err != nil {
		return
	}
	uid = user.Id
	return
}

func (u *user) SetPhone(uid int, code string) (err error) {
	res, err := u.getPhoneNumber(code)
	if err != nil {
		return
	}
	if err = res.GetResponseError(); err != nil {
		return
	}
	user := model.User{Id: uid}
	user.Phone = res.Data.PhoneNumber
	err = user.Update()
	return
}

func (u *user) SetNameAndAvatarUrl(uid int, nickName string, avatarUrl string) (err error) {
	user := model.User{Id: uid}
	user.NickName = nickName
	user.AvatarUrl = avatarUrl
	err = user.Update()
	return
}

func (u *user) BackendLogin(account string, password string, ip string) (token string, err error) {
	for _, bgAccountInfo := range *conf.App.BgAccounts {
		if account != bgAccountInfo.Account {
			continue
		}
		if password != bgAccountInfo.Password {
			break
		}
		token, err = utils.CreateAuthToken(bgAccountInfo.Uid, ip)
		return
	}
	err = errors.New("帐号或密码不正确。")
	return
}

func (u *user) wxLogin(code string) (*weapp.LoginResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	return sdk.Login(code)
}

func (u *user) getPhoneNumber(code string) (*phonenumber.GetPhoneNumberResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	cli := sdk.NewPhonenumber()
	return cli.GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: code})
}
