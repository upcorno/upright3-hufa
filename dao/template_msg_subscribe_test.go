package dao

import (
	"errors"
	"testing"
)

var templateId string = "sssss"

func TestTemplateMsgSubscribe(t *testing.T) {
	TMsgSubDao.insert(TestUid, templateId, 0)
	defer TMsgSubDao.delete(TestUid, templateId)
	err := TMsgSubDao.IncrSubscribeNum(TestUid, templateId)
	if err != nil {
		t.Fatal(err)
	}
	subNum, err := TMsgSubDao.getSubscribeNum(TestUid, templateId)
	if err != nil {
		t.Fatal(err)
	}
	if subNum != 1 {
		t.Fatal(errors.New("此时订阅次数应为1"))
	}
	err = TMsgSubDao.IncrSubscribeNum(TestUid, templateId)
	if err != nil {
		t.Fatal(err)
	}
	err = TMsgSubDao.DecrSubscribeNum(TestUid, templateId)
	if err != nil {
		t.Fatal(err)
	}
	subNum, err = TMsgSubDao.getSubscribeNum(TestUid, templateId)
	if err != nil {
		t.Fatal(err)
	}
	if subNum != 1 {
		t.Fatal(errors.New("此时订阅次数应为1"))
	}
	//即使subNum=0，多次删除不应报错
	TMsgSubDao.DecrSubscribeNum(TestUid, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUid, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUid, templateId)
	TMsgSubDao.DecrSubscribeNum(TestUid, templateId)
}
