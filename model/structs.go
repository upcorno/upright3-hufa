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
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question"`
	Imgs          string    `xorm:"comment('描述图片') TEXT" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT(10)" json:"consultant_uid"`
	Status        string    `xorm:"not null default '处理中' comment('处理中，待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status"`
	CreateTime    int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type ConsultationRecord struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT(10)" json:"id"`
	ConsultationId  int       `xorm:"not null comment('咨询id') index UNSIGNED INT(10)" json:"consultation_id"`
	CommunicatorUid int       `xorm:"not null comment('沟通人uid') UNSIGNED INT(10)" json:"communicator_uid"`
	Type            string    `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type"`
	Content         string    `xorm:"comment('回复内容') LONGTEXT" json:"content"`
	CreateTime      int       `xorm:"INT(10)" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime" json:"update_time"`
}

type Favorites struct {
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
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"update_time"`
}
