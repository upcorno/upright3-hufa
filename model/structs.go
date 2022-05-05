package model

import (
	"time"
)

//常见知产问题
type LegalIssue struct {
	Id             int       `xorm:"pk autoincr INT" json:"id"`
	CreatorUid     int       `xorm:"not null comment('问题创建人id') index UNSIGNED INT" json:"creator_uid"`
	FirstCategory  string    `xorm:"not null comment('一级类别') index CHAR(6)" json:"first_category"`
	SecondCategory string    `xorm:"not null comment('二级类别') index CHAR(25)" json:"second_category"`
	Tags           string    `xorm:"comment('问题标签') index VARCHAR(255)" json:"tags"`
	Title          string    `xorm:"not null comment('标题') VARCHAR(60)" json:"title"`
	Imgs           string    `xorm:"comment('普法问题关联图片') TEXT" json:"imgs"`
	Content        string    `xorm:"comment('内容') LONGTEXT" json:"content"`
	SearchText     string    `xorm:"comment('全文检索字段') LONGTEXT" json:"-"`
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

// 问题“咨询”
type Consultation struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question" validate:"required"`
	Imgs          string    `xorm:"comment('描述图片') TEXT" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT" json:"consultant_uid" validate:"required"`
	Status        string    `xorm:"not null default '处理中' comment('处理中、待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status" validate:"required,oneof=处理中 待人工咨询 人工咨询中 已完成"`
	CreateTime    int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

// “咨询”沟通记录
type ConsultationReply struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	ConsultationId  int       `xorm:"not null comment('咨询id') index UNSIGNED INT" json:"consultation_id"`
	CommunicatorUid int       `xorm:"not null comment('沟通人uid') UNSIGNED INT" json:"communicator_uid"`
	Type            string    `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type" validate:"required,oneof=answer query"`
	Content         string    `xorm:"comment('回复内容') LONGTEXT" json:"content" validate:"required"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//“侵权监测”提交信息
type InfringementMonitor struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
	Organization    string    `xorm:"comment('组织机构') VARCHAR(60)" json:"organization"`
	Description     string    `xorm:"comment('侵权监测描述') TEXT" json:"description"`
	Resume          string    `xorm:"comment('权利概要') TEXT" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50)" json:"customer_address"`
	DealRemark      string    `xorm:"comment('回访时备注') TEXT" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//“我要维权”用户提交信息
type RightsProtection struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
	Organization    string    `xorm:"comment('组织结构') VARCHAR(60)" json:"organization"`
	Description     string    `xorm:"comment('维权意向描述') TEXT" json:"description"`
	Resume          string    `xorm:"comment('权利概要') TEXT" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50)" json:"customer_address"`
	DealRemark      string    `xorm:"comment('回访时备注') TEXT" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type User struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	AppId      string    `xorm:"not null comment('公众号/小程序标识') unique(app_id) CHAR(20)" json:"app_id"`
	Openid     string    `xorm:"not null comment('与appid共同锁定一个用户') unique(app_id) CHAR(30)" json:"openid"`
	Unionid    string    `xorm:"comment('同一开放平台下，各公众号/小程序具有相同unionid') index CHAR(30)" json:"unionid"`
	NickName   string    `xorm:"comment('用户微信昵称') VARCHAR(16)" json:"nick_name"`
	AvatarUrl  string    `xorm:"comment('头像地址') TEXT" json:"avatar_url"`
	Phone      string    `xorm:"comment('电话号码') CHAR(20)" json:"phone"`
	CreateTime int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}
