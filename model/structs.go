package model

import (
	"time"
)

type Category struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	Name          string    `xorm:"not null comment('类别名') VARCHAR(20)" json:"name"`
	PreCategoryId int       `xorm:"comment('上一级分类id') index UNSIGNED INT(10)" json:"pre_category_id"`
	CreateTime    int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"update_time"`
}

type Consultation struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question" validate:"required"`
	Imgs          string    `xorm:"comment('描述图片') TEXT" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT(10)" json:"consultant_uid" validate:"required"`
	Status        string    `xorm:"not null default '处理中' comment('处理中，待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status" validate:"required,oneof=处理中 待人工咨询 人工咨询中 已完成"`
	CreateTime    int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type ConsultationRecord struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	ConsultationId  int       `xorm:"not null comment('咨询id') index UNSIGNED INT(10)" json:"consultation_id"`
	CommunicatorUid int       `xorm:"not null comment('沟通人uid') UNSIGNED INT(10)" json:"communicator_uid"`
	Type            string    `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type" validate:"required,oneof=answer query"`
	Content         string    `xorm:"comment('回复内容') LONGTEXT" json:"content" validate:"required"`
	CreateTime      int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type Favorite struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	UserId     int       `xorm:"not null comment('用户id') index UNSIGNED INT(10)" json:"user_id"`
	IssueId    int       `xorm:"not null comment('普法知识问题id') index UNSIGNED INT(10)" json:"issue_id"`
	CreateTime int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type LegalIssue struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	CreatorUid int       `xorm:"not null comment('问题创建人id') index UNSIGNED INT(10)" json:"creator_uid"`
	CategoryId int       `xorm:"not null comment('类别id') index UNSIGNED INT(10)" json:"category_id"`
	Title      string    `xorm:"not null comment('标题') index index VARCHAR(60)" json:"title"`
	Imgs       string    `xorm:"comment('普法问题关联图片') TEXT" json:"imgs"`
	IssueText  string    `xorm:"comment('内容') LONGTEXT" json:"issue_text"`
	CreateTime int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"update_time"`
}

type User struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	AppId      string    `xorm:"not null comment('公众号/小程序标识') unique(app_id) CHAR(20)" json:"app_id"`
	Openid     string    `xorm:"not null comment('与appid共同锁定一个用户') unique(app_id) CHAR(30)" json:"openid"`
	Unionid    string    `xorm:"comment('同一开放平台下，各公众号/小程序具有相同unionid') index CHAR(30)" json:"unionid"`
	NickName   string    `xorm:"comment('用户微信昵称') VARCHAR(16)" json:"nick_name"`
	AvatarUrl  string    `xorm:"comment('头像地址') TEXT" json:"avatar_url"`
	Phone      string    `xorm:"comment('电话号码') CHAR(20)" json:"phone"`
	AddTime    int       `xorm:"INT(10)" json:"add_time"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" json:"update_time"`
}

type InfringementDetection struct {
	Id           int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	CreatorUid   int       `xorm:"not null comment('创建人id') index UNSIGNED INT(10)" json:"creator_uid"`
	Name         string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone        string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
	Organization string    `xorm:"comment('组织机构') VARCHAR(60)" json:"organization"`
	Description  string    `xorm:"comment('侵权监测描述') TEXT" json:"description"`
	Resume       string    `xorm:"comment('权利概要') TEXT" json:"resume"`
	CreateTime   int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime   time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type DetectionReturnVisit struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	CreatorUid      int       `xorm:"not null comment('创建人id') index UNSIGNED INT(10)" json:"creator_uid" validate:"required"`
	DetectionId     int       `xorm:"not null comment('监测id') UNSIGNED INT(10)" json:"detection_id" validate:"required"`
	Classification  string    `xorm:"not null comment('类别') VARCHAR(10)" json:"classification" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string    `xorm:"not null comment('客户地址') VARCHAR(50)" json:"customer_address"`
	Remark          string    `xorm:"comment('备注') TEXT" json:"remark"`
	CreateTime      int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime" json:"update_time"`
}
type RightsProtection struct {
	Id           int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	CreatorUid   int       `xorm:"not null comment('创建人id') index UNSIGNED INT(10)" json:"creator_uid"`
	Name         string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name" validate:"required"`
	Phone        string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required"`
	Organization string    `xorm:"comment('组织结构') VARCHAR(60)" json:"organization"`
	Description  string    `xorm:"comment('维权意向描述') TEXT" json:"description"`
	Resume       string    `xorm:"comment('权利概要') TEXT" json:"resume"`
	CreateTime   int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime   time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type ProtectionReturnVisit struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	CreatorUid      int       `xorm:"not null comment('创建人id') index UNSIGNED INT(10)" json:"creator_uid" validate:"required"`
	ProtectionId    int       `xorm:"not null comment('维权id') UNSIGNED INT(10)" json:"protection_id" validate:"required"`
	Classification  string    `xorm:"not null comment('类别') VARCHAR(10)" json:"classification" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string    `xorm:"not null comment('客户地址') VARCHAR(50)" json:"customer_address"`
	Remark          string    `xorm:"comment('备注') TEXT" json:"remark"`
	CreateTime      int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime" json:"update_time"`
}
