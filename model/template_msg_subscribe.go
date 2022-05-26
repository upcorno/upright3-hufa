package model

import (
	"fmt"
	"time"
)

//微信模版消息订阅记录
type TemplateMsgSubscribe struct {
	Id           int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	UserId       int       `xorm:"not null UNSIGNED INT unique(user_id_template_id)" json:"user_id"`
	TemplateId   string    `xorm:"not null comment('微信模版消息id') index CHAR(60) unique(user_id_template_id)" json:"template_id"`
	SubscribeNum int       `xorm:"not null comment('订阅次数') UNSIGNED INT default 0" json:"subscribe_num"`
	UpdateTime   time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type templateMsgSubscribeDao struct{}

var TMsgSubDao *templateMsgSubscribeDao

func (t *templateMsgSubscribeDao) insert(userId int, templateId string, subscribeNum int) (sub *TemplateMsgSubscribe, err error) {
	if !(userId > 0 && len(templateId) < 60) {
		err = fmt.Errorf("userId应大于0，且templateId长度不高于60.userId:%d,templateId:%s", userId, templateId)
		return
	}
	sub = &TemplateMsgSubscribe{
		UserId:       userId,
		TemplateId:   templateId,
		SubscribeNum: subscribeNum,
	}
	_, err = Db.InsertOne(sub)
	return
}

func (t *templateMsgSubscribeDao) delete(userId int, templateId string) (err error) {
	sub := TemplateMsgSubscribe{
		UserId:       userId,
		TemplateId:   templateId,
		SubscribeNum: 0,
	}
	_, err = Db.Delete(sub)
	return
}

// 查询订阅次数
func (t *templateMsgSubscribeDao) getSubscribeNum(userId int, templateId string) (subscribeNum int, err error) {
	bean := TemplateMsgSubscribe{
		UserId:     userId,
		TemplateId: templateId,
	}
	has, err := Db.Get(&bean)
	if err != nil {
		return
	}
	if !has {
		err = fmt.Errorf("未查询到用户模版消息订阅记录.userId:%d,templateId:%s", userId, templateId)
		return
	}
	subscribeNum = bean.SubscribeNum
	return
}

// 增加订阅次数
func (t *templateMsgSubscribeDao) IncrSubscribeNum(userId int, templateId string) (err error) {
	session := Db.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		session.Rollback()
		return
	}
	sub := &TemplateMsgSubscribe{
		UserId:     userId,
		TemplateId: templateId,
	}
	has, err := session.ForUpdate().Get(sub)
	if err != nil {
		session.Rollback()
		return
	}
	if !has {
		_, err = t.insert(userId, templateId, 1)
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
		return
	}
	_, err = session.Incr("subscribe_num", 1).Update(&TemplateMsgSubscribe{}, &TemplateMsgSubscribe{
		UserId:     userId,
		TemplateId: templateId,
	})
	if err != nil {
		session.Rollback()
		return
	}
	session.Commit()
	return
}

//DecrSubscribeNum 减少订阅次数
func (t *templateMsgSubscribeDao) DecrSubscribeNum(userId int, templateId string) (err error) {
	session := Db.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		session.Rollback()
		return
	}
	bean := &TemplateMsgSubscribe{
		UserId:     userId,
		TemplateId: templateId,
	}
	has, err := session.ForUpdate().Get(bean)
	if err != nil {
		session.Rollback()
		return
	}
	if !has {
		err = fmt.Errorf("未查询到用户模版消息订阅记录.userId:%d,templateId:%s", userId, templateId)
		session.Rollback()
		return
	}
	if bean.SubscribeNum < 1 {
		session.Rollback()
		return
	}
	_, err = session.Decr("subscribe_num", 1).Update(&TemplateMsgSubscribe{}, &TemplateMsgSubscribe{
		UserId:     userId,
		TemplateId: templateId,
	})
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Commit()
	return
}
