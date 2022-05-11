package model

import (
	"errors"
	"time"
)

// “咨询”沟通记录
type ConsultationReply struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	ConsultationId  int       `xorm:"not null comment('咨询id') index UNSIGNED INT" json:"consultation_id"`
	CommunicatorUid int       `xorm:"not null comment('沟通人uid') UNSIGNED INT" json:"communicator_uid"`
	Type            string    `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type"`
	Content         string    `xorm:"not null comment('回复内容') LONGTEXT" json:"content"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//创建咨询回复记录
func (reply *ConsultationReply) Insert() (err error) {
	if reply.ConsultationId == 0 || reply.CommunicatorUid == 0 || reply.Type == "" || reply.Content == "" {
		err = errors.New("model-ConsultationReply:ConsultationId、CommunicatorUid、Type、Content cannot be empty")
		return
	}
	reply.CreateTime = int(time.Now().Unix())
	_, err = Db.InsertOne(reply)
	return
}

//删除咨询回复记录
func (reply *ConsultationReply) delete() (err error) {
	if reply.Id == 0 {
		err = errors.New("model:必须指定id值")
		return
	}
	_, err = Db.Delete(reply)
	return
}

type replyInfo struct {
	Id              int    `json:"id"`
	ConsultationId  int    `json:"consultation_id"`
	CommunicatorUid int    `json:"communicator_uid"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	NickName        string `json:"nick_name"`
	AvatarUrl       string `json:"avatar_url"`
	Phone           string `json:"phone"`
	CreateTime      int    `json:"create_time"`
}

//获取咨询沟通记录表
func ConsultationReplyList(consultationId int) (replyInfoList []replyInfo, err error) {
	if consultationId == 0 {
		err = errors.New("model:必须指定consultationId值")
		return
	}
	err = Db.Table("consultation_reply").
		Join("INNER", "user", "user.id = consultation_reply.communicator_uid").
		Where("consultation_id=?", consultationId).
		Cols(
			"consultation_reply.id",
			"consultation_id",
			"communicator_uid",
			"type",
			"content",
			"nick_name",
			"avatar_url",
			"phone",
			"consultation_reply.create_time",
		).
		Asc("consultation_reply.create_time").
		Find(&replyInfoList)
	return
}
