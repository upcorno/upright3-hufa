package service

import (
	"errors"
	"law/conf"
	dao "law/dao"
	"law/utils"
)

type user struct{}

var UserSrv *user = &user{}

func (u *user) Login(code string) (token string, err error) {
	openid, unionid, err := WxSrv.wxLogin(code)
	if err != nil {
		return
	}
	uid, err := u.sync(openid, unionid)
	if err != nil {
		return
	}
	return utils.CreateAuthToken(uid)
}

// 存储openid、unionid，并返回userId
func (u *user) sync(openid string, unionid string) (userId int, err error) {
	has, user, err := dao.UserDao.Get(0, conf.App.WxApp.Appid, openid)
	if err != nil {
		return
	}
	if has {
		userId = user.Id
		if unionid != "" && user.Unionid != unionid {
			//更新 unionid, unionid 可能从 null 变为由内容
			user.Unionid = unionid
			err = dao.UserDao.Update(user.Id, user)
		}
		return
	}
	user = &dao.User{
		AppId:   conf.App.WxApp.Appid,
		Openid:  openid,
		Unionid: unionid,
	}
	if userId, err = dao.UserDao.Insert(user); err != nil {
		return
	}
	return
}

func (u *user) SetPhone(userId int, code string) (err error) {
	phoneNumber, err := WxSrv.getPhoneNumber(code)
	if err != nil {
		return
	}
	err = dao.UserDao.Update(userId, &dao.User{Phone: phoneNumber})
	return
}

func (u *user) SetNameAndAvatarUrl(userId int, nickName string, avatarUrl string) (err error) {
	user := &dao.User{NickName: nickName, AvatarUrl: avatarUrl}
	err = dao.UserDao.Update(userId, user)
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
