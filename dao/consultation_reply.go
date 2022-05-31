package dao

import (
	"errors"
	"law/enum"
	"time"
)

// “咨询”沟通记录
type ConsultationReply struct {
	Id              int          `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	ConsultationId  int          `xorm:"not null comment('咨询id') index UNSIGNED INT" json:"consultation_id"`
	CommunicatorUid int          `xorm:"not null comment('沟通人uid') UNSIGNED INT" json:"communicator_uid"`
	Type            string       `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type"`
	Content         string       `xorm:"not null comment('回复内容') LONGTEXT" json:"content"`
	IsMannerReply   enum.YesOrNo `xorm:"not null comment('是否为人工回复') CHAR(3) default('no')" json:"is_manner_reply"`
	CreateTime      int          `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time    `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type consulReplyDao struct{}

var ConsulReplyDao *consulReplyDao

//创建咨询回复记录
func (c *consulReplyDao) Insert(consulId int, category string, content string, communicatorUid int, isMannerReply enum.YesOrNo) (replyId int, err error) {
	reply := &ConsultationReply{
		ConsultationId:  consulId,
		Type:            category,
		Content:         content,
		CommunicatorUid: communicatorUid,
		IsMannerReply:   isMannerReply,
		CreateTime:      int(time.Now().Unix()),
	}
	if reply.ConsultationId == 0 || reply.CommunicatorUid == 0 || reply.Type == "" || reply.Content == "" {
		err = errors.New("dao-ConsultationReply:ConsultationId、CommunicatorUid、Type、Content cannot be empty")
		return
	}
	_, err = Db.InsertOne(reply)
	if err == nil {
		replyId = reply.Id
	}
	return
}

//删除咨询回复记录
func (c *consulReplyDao) delete(replyId int) (err error) {
	if replyId == 0 {
		err = errors.New("replyId不可为0")
		return
	}
	_, err = Db.Delete(&ConsultationReply{Id: replyId})
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
func (c *consulReplyDao) List(consulId int, minReplyId int, onlyMannerReply bool) (replyInfoList []replyInfo, err error) {
	if consulId == 0 {
		err = errors.New("dao:必须指定consultationId值")
		return
	}
	replyInfoList = []replyInfo{}
	sess := Db.Table("consultation_reply").
		Join("INNER", "user", "user.id = consultation_reply.communicator_uid").
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
		Asc("consultation_reply.create_time")
	sess.Where("consultation_id=?", consulId).
		And("consultation_reply.id>?", minReplyId)
	if onlyMannerReply {
		sess.And("is_manner_reply=?", enum.YES)
	}
	err = sess.Find(&replyInfoList)
	return
}
