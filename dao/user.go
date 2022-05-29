package dao

import (
	"errors"
	"time"
)

type User struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	AppId      string    `xorm:"not null comment('公众号/小程序标识') unique(app_id) CHAR(20)" json:"app_id"`
	Openid     string    `xorm:"not null comment('与appid共同锁定一个用户') unique(app_id) CHAR(30)" json:"openid"`
	Unionid    string    `xorm:"not null comment('同一开放平台下，各公众号/小程序具有相同unionid') index CHAR(30) default('')" json:"unionid"`
	NickName   string    `xorm:"not null comment('用户微信昵称') VARCHAR(16) default('')" json:"nick_name"`
	AvatarUrl  string    `xorm:"not null comment('头像地址') TEXT default('')" json:"avatar_url"`
	Phone      string    `xorm:"not null comment('电话号码') CHAR(20) default('')" json:"phone"`
	CreateTime int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type userDao struct{}

var UserDao *userDao

func (u *userDao) Insert(user *User) (userId int, err error) {
	if user.AppId == "" || user.Openid == "" {
		err = errors.New("AppId、Openid不可以为空值")
		return
	}
	if user.CreateTime == 0 {
		user.CreateTime = int(time.Now().Unix())
	}
	_, err = Db.InsertOne(user)
	if err == nil {
		userId = user.Id
	}
	return
}

func (u *userDao) Get(userId int, appId string, openid string) (has bool, user *User, err error) {
	user = &User{
		Id:     userId,
		AppId:  appId,
		Openid: openid,
	}
	if user.Id == 0 {
		if !(user.AppId != "" && user.Openid != "") {
			err = errors.New("dao:查询用户时须指定id值或通过appid、openid获取。")
			return
		}
	}
	has, err = Db.Get(user)
	return
}

func (u *userDao) delete(userId int, appId string, openid string) (err error) {
	user := &User{
		Id:     userId,
		AppId:  appId,
		Openid: openid,
	}
	if user.Id == 0 {
		if !(user.AppId != "" && user.Openid != "") {
			err = errors.New("dao:查询用户时须指定id值或通过appid、openid获取。")
			return
		}
	}
	_, err = Db.Delete(user)
	return
}

func (u *userDao) Update(userId int, user *User) (err error) {
	if userId == 0 {
		err = errors.New("userId不可为0值")
		return
	}
	_, err = Db.Update(user, User{Id: userId})
	return
}
