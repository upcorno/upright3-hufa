package model

import (
	"time"
)

//常见知产问题
type LegalIssue struct {
	Id             int       `xorm:"not null pk autoincr INT" json:"id"`
	CreatorUid     int       `xorm:"not null comment('问题创建人id') index UNSIGNED INT" json:"creator_uid"`
	FirstCategory  string    `xorm:"not null comment('一级类别') index CHAR(6)" json:"first_category"`
	SecondCategory string    `xorm:"not null comment('二级类别') index CHAR(25)" json:"second_category"`
	Tags           string    `xorm:"not null comment('问题标签') index VARCHAR(255) default('')" json:"tags"`
	Title          string    `xorm:"not null comment('标题') VARCHAR(60)" json:"title"`
	Imgs           string    `xorm:"not null comment('普法问题关联图片') TEXT default('')" json:"imgs"`
	Content        string    `xorm:"not null comment('内容') LONGTEXT" json:"content"`
	SearchText     string    `xorm:"not null comment('全文检索字段') LONGTEXT default('')" json:"-"`
	CreateTime     int       `xorm:"not null UNSIGNED INT default(1651383059)" json:"create_time"`
	UpdateTime     time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//用户收藏
type LegalIssueFavorite struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	UserId     int       `xorm:"not null comment('用户id') index UNSIGNED INT" json:"user_id"`
	IssueId    int       `xorm:"not null comment('普法知识问题id') index UNSIGNED INT" json:"issue_id"`
	CreateTime int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

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

//“我要维权”用户提交信息
type RightsProtection struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
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
