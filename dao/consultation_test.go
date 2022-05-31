package dao

import (
	"errors"
	"law/enum"
	"testing"
)

func TestConsultation(t *testing.T) {
	consulId, err := ConsulDao.Insert("这是测试的问题", "", TestUserId, enum.DOING)
	if err != nil {
		t.Fatal(err)
	}
	if consulId == 0 {
		t.Fatal(errors.New("添加consul失败"))
	}
	consultationList, err := ConsulDao.List(TestUserId)
	if err != nil {
		t.Fatal(err)
	}
	if len(consultationList) != 1 {
		t.Fatal(errors.New("consultationList应包含一个consultation"))
	}
	testReply(consulId, t)
	testSetStatus(consulId, t)
	testConsultationGetWithUserInfo(consulId, t)
	testGetUnexistConsultation(consulId+1, t)
	err = ConsulDao.delete(consulId)
	if err != nil {
		t.Fatal(err)
	}
}

func testReply(consulId int, t *testing.T) {
	replyId, err := ConsulReplyDao.Insert(consulId, "answer", "单元测试", TestUserId, enum.YES)
	if err != nil {
		t.Fatal(err)
	}
	replyInfoList, err := ConsulReplyDao.List(consulId, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(replyInfoList) != 1 {
		t.Fatal(errors.New("Consultation应存在回复"))
	}
	ConsulReplyDao.delete(replyId)

	replyInfoList, err = ConsulReplyDao.List(consulId, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(replyInfoList) != 0 {
		t.Fatal(errors.New("Consultation不应存在回复"))
	}
}

func testSetStatus(consultationId int, t *testing.T) {
	consul := &Consultation{Status: enum.DONE}
	ConsulDao.Update(consultationId, consul, "status")
	consulWithUserInfo, err := ConsulDao.GetWithUserInfo(consultationId)
	if err != nil {
		t.Fatal(err)
	}
	if consulWithUserInfo == nil || consulWithUserInfo.Status != consul.Status {
		t.Fatal("修改Consultation状态失败")
	}
}

func testConsultationGetWithUserInfo(consultationId int, t *testing.T) {
	consulWithUserInfo, err := ConsulDao.GetWithUserInfo(consultationId)
	if err != nil {
		t.Fatal(err)
	}
	if consulWithUserInfo == nil {
		t.Fatal("查询ConsultationGetWithUser失败。")
	}
}
func testGetUnexistConsultation(unexixtConsultationId int, t *testing.T) {
	consulPlus, err := ConsulDao.GetWithUserInfo(unexixtConsultationId)
	if err != nil {
		t.Fatal(err)
	}
	if consulPlus != nil {
		t.Fatal(errors.New("该Consultation不应存在"))
	}
}
