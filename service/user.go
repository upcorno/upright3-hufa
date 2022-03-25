package service

import (
	"fmt"
	"law/conf"
	"law/model"

	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

func Login(code string) (uid int, err error) {
	res, err := wxLogin(code)
	if err != nil {
		return
	}
	if err = res.GetResponseError(); err != nil {
		return
	}
	return getUid(res.OpenID)
}

// 根据 openId 获取用户 id，不存在时创建新用户返回对应 id
func getUid(openid string) (uid int, err error) {
	user := model.User{Openid: openid}
	has, err := user.Get()
	if err != nil {
		return
	}
	if has {
		uid = user.Id
		return
	}
	user.AppId = conf.App.WxApp.Appid
	if err = user.Insert(); err != nil {
		return
	}
	uid = user.Id
	return
}

func SetPhone(uid int, code string) (err error) {
	res, err := getPhoneNumber(code)
	if err != nil {
		return
	}
	if err = res.GetResponseError(); err != nil {
		return
	}
	user := model.User{Id: uid}
	has, err := user.Get()
	if err != nil {
		return
	}
	if !has {
		return fmt.Errorf("用户 uid(%d) 不存在", uid)
	}
	user.Phone = res.Data.PhoneNumber
	if err = user.Update(); err != nil {
		return
	}
	return
}

func SetNameAndAvatarUrl(uid int, nickName string, avatarUrl string) (err error) {
	user := model.User{Id: uid}
	has, err := user.Get()
	if err != nil {
		return
	}
	if !has {
		return fmt.Errorf("用户 uid(%d) 不存在", uid)
	}
	user.NickName = nickName
	user.AvatarUrl = avatarUrl
	if err = user.Update(); err != nil {
		return
	}
	return
}

func wxLogin(code string) (*weapp.LoginResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	return sdk.Login(code)
}

func getPhoneNumber(code string) (*phonenumber.GetPhoneNumberResponse, error) {
	sdk := weapp.NewClient(conf.App.WxApp.Appid, conf.App.WxApp.Secret)
	cli := sdk.NewPhonenumber()
	return cli.GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: code})
}
