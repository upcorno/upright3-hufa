package model

import (
	"time"
)

//“侵权监测”提交信息
type InfringementMonitor struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
	Organization    string    `xorm:"not null comment('组织机构') VARCHAR(60) default('')" json:"organization"`
	Description     string    `xorm:"not null comment('侵权监测描述') TEXT default('')" json:"description"`
	Resume          string    `xorm:"not null comment('权利概要') TEXT default('')" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50) default('')" json:"customer_address"`
	DealRemark      string    `xorm:"not null comment('回访时备注') TEXT default('')" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null unique comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}
