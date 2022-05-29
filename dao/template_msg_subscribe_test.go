package dao

import (
	"errors"
	"testing"
)

var templateId string = "sssss"

func TestTemplateMsgSubscribe(t *testing.T) {
	TMsgSubDao.insert(TestUserId, templateId, 0)
	defer TMsgSubDao.delete(TestUserId, templateId)
	err := TMsgSubDao.IncrSubscribeNum(TestUserId, templateId)
	if err != nil {
		t.Fatal(err)
	}
	subNum, err := TMsgSubDao.getSubscribeNum(TestUserId, templateId)
	if err != nil {
		t.Fatal(err)
	}
	if subNum != 1 {
		t.Fatal(errors.New("此时订阅次数应为1"))
	}
	err = TMsgSubDao.IncrSubscribeNum(TestUserId, templateId)
	if err != nil {
		t.Fatal(err)
	}
	err = TMsgSubDao.DecrSubscribeNum(TestUserId, templateId)
	if err != nil {
		t.Fatal(err)
	}
	subNum, err = TMsgSubDao.getSubscribeNum(TestUserId, templateId)
	if err != nil {
		t.Fatal(err)
	}
	if subNum != 1 {
		t.Fatal(errors.New("此时订阅次数应为1"))
	}
	//即使subNum=0，多次删除不应报错
	TMsgSubDao.DecrSubscribeNum(TestUserId, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUserId, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUserId, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUserId, templateId)
}
