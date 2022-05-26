package service

import (
	"errors"
	"law/conf"
	"law/model"
	"law/utils"
	"time"
)

type user struct{}

var UserSrv *user = &user{}

func (u *user) Login(code string) (token string, err error) {
	openid, unionid, err := WxSrv.wxLogin(code)
	if err != nil {
		return
	}
	uid, err := u.getUidAndSync(openid, unionid)
	if err != nil {
		return
	}
	return utils.CreateAuthToken(uid)
}

// 根据 openId 获取用户 id，不存在时创建新用户返回对应 id
func (u *user) getUidAndSync(openid string, unionid string) (uid int, err error) {
	user := &model.User{
		AppId:  conf.App.WxApp.Appid,
		Openid: openid,
	}
	has, err := user.Get()
	if err != nil {
		return
	}
	if has {
		uid = user.Id
		if unionid != "" && user.Unionid != unionid {
			//更新 unionid, unionid 可能从 null 变为由内容
			err = user.Update()
		}
		return
	}
	user.Unionid = unionid
	user.CreateTime = int(time.Now().Unix())
	if err = user.Insert(); err != nil {
		return
	}
	uid = user.Id
	return
}

func (u *user) SetPhone(uid int, code string) (err error) {
	phoneNumber, err := WxSrv.getPhoneNumber(code)
	if err != nil {
		return
	}
	user := model.User{Id: uid}
	user.Phone = phoneNumber
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

func (u *user) BackendLogin(account string, password string) (token string, err error) {
	for _, bgAccountInfo := range *conf.App.BgAccounts {
		if account != bgAccountInfo.Account {
			continue
		}
		if password != bgAccountInfo.Password {
			break
		}
		token, err = utils.CreateAuthToken(bgAccountInfo.Uid)
		return
	}
	err = errors.New("帐号或密码不正确。")
	return
}
