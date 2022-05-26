package dao

import (
	"errors"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

//“我要监测”用户提交信息
//CreatorUid字段具有唯一性约束。即一个用户只能有一个
type InfringementMonitor struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone"`
	Organization    string    `xorm:"not null comment('组织结构') VARCHAR(60) default('')" json:"organization"`
	Description     string    `xorm:"not null comment('维权意向描述') TEXT default('')" json:"description"`
	Resume          string    `xorm:"not null comment('权利概要') TEXT default('')" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50) default('')" json:"customer_address"`
	DealRemark      string    `xorm:"comment('回访时备注') TEXT default('')" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null unique comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

func (r *InfringementMonitor) Insert() (err error) {
	if r.Name == "" || r.Phone == "" || r.CreatorUid == 0 {
		err = errors.New("必须指定Name、Phone、CreatorUid字段")
		return
	}
	if utf8.RuneCountInString(r.Name) > 16 || utf8.RuneCountInString(r.Phone) > 20 {
		err = errors.New("Name不可超过16个字符、且Phone不超过20个字符")
		return
	}
	r.CreateTime = int(time.Now().Unix())
	_, err = Db.InsertOne(r)
	return
}

func (r *InfringementMonitor) Get() (has bool, err error) {
	if r.Id == 0 {
		if r.CreatorUid == 0 {
			err = errors.New("需指定InfringementMonitor的Id或CreatorUid字段")
			return
		}
	}
	has, err = Db.Get(r)
	return
}

func (r *InfringementMonitor) Update(columns ...string) (err error) {
	if r.Id == 0 {
		if r.CreatorUid == 0 {
			err = errors.New("需指定InfringementMonitor的Id或CreatorUid字段")
			return
		}
	}
	_, err = Db.Cols(columns...).Update(r, &InfringementMonitor{Id: r.Id, CreatorUid: r.CreatorUid})
	return
}

func (r *InfringementMonitor) delete() (err error) {
	if r.Id == 0 {
		if r.CreatorUid == 0 {
			err = errors.New("需指定InfringementMonitor的Id或CreatorUid字段")
			return
		}
	}
	_, err = Db.Delete(&InfringementMonitor{Id: r.Id, CreatorUid: r.CreatorUid})
	return
}
