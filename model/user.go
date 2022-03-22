package model

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	YsjUid           int       `xorm:"pk INT(11)" json:"ysj_uid"`
	Roles            []string  `xorm:"comment('用户角色') LONGTEXT" json:"roles"`
	Name             string    `xorm:"not null VARCHAR(255)" json:"name"`
	AvatarUrl        string    `xorm:"VARCHAR(255)" json:"avatar_url"`
	JwtLoginTokenExp int       `xorm:"INT(11)" json:"-"` //登陆令牌过期时间,保证登陆令牌仅使用一次
	CreateTime       time.Time `xorm:"not null created DateTime" json:"create_time"`
	UpdateTime       time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

func UserCreate(mod *User) error {
	_, err := Db.InsertOne(mod)
	return err
}

func UserGetByYsjUid(ysjUid int) (*User, bool, error) {
	mod := &User{}
	has, err := Db.Where("ysj_uid=?", ysjUid).Cols("jwt_login_token_exp").Get(mod)
	return mod, has, err
}

func UserUpdateByYsjUid(ysjUid int, mod *User) error {
	_, error := Db.Where("ysj_uid=?", ysjUid).Update(mod)
	return error
}

func UserExistCheckByYsjUid(ysjUid int) (bool, error) {
	mod := &User{
		YsjUid: ysjUid,
	}
	has, err := Db.Cols("uid").Get(mod)
	if err != nil {
		return false, err
	}
	return has, nil
}

func UsersList(role string) ([]User, error) {
	users := make([]User, 0)
	err := Db.Where("roles like ?", "%"+role+"%").Find(&users)
	return users, err
}

func AllUsersList() ([]User, error) {
	users := make([]User, 0)
	err := Db.Where("1=1").Find(&users)
	return users, err
}

func UserInfo(uid int) (*User, error) {
	users := make([]User, 0)
	err := Db.Where("ysj_uid = ?", uid).Limit(1).Find(&users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user not found.uid:%d", uid)
	}
	return &users[0], err
}

func CurrentUserUCurrencyAmount(uid int) ([]int, error) {
	var ints []int
	err := Db1.Table("ucurrency_detail").Cols("u_currency_amount").Where("type = ? and beneficiary_uid = ?", "反馈奖励", uid).Find(&ints)
	return ints, err
}
