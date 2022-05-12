package model

import (
	"errors"
	"law/enum"
	"testing"
	"time"
)

func TestConsultation(t *testing.T) {
	consul := &Consultation{
		Question:      "这是测试的问题",
		Status:        enum.DOING,
		ConsultantUid: TestUid,
		CreateTime:    int(time.Now().Unix()),
	}
	err := consul.Create()
	if err != nil {
		t.Fatal(err)
	}
	if consul.Id == 0 {
		t.Fatal(errors.New("添加consul失败"))
	}
	consultationList, err := ConsultationList(TestUid)
	if err != nil {
		t.Fatal(err)
	}
	if len(consultationList) != 1 {
		t.Fatal(errors.New("consultationList应包含一个consultation"))
	}
	testReply(consul.Id, t)
	testSetStatus(consul.Id, t)
	testConsultationGetWithUserInfo(consul.Id, t)
	testGetUnexixtConsultation(consul.Id+1, t)
	err = consul.delete()
	if err != nil {
		t.Fatal(err)
	}
}

func testReply(consultationId int, t *testing.T) {
	reply := &ConsultationReply{
		ConsultationId:  consultationId,
		CommunicatorUid: TestUid,
		Type:            "answer",
		Content:         "单元测试",
		CreateTime:      int(time.Now().Unix()),
	}
	err := reply.Insert()
	if err != nil {
		t.Fatal(err)
	}
	replyInfoList, err := ConsultationReplyList(consultationId)
	if err != nil {
		t.Fatal(err)
	}
	if len(replyInfoList) != 1 {
		t.Fatal(errors.New("Consultation应存在回复"))
	}
	reply.delete()

	replyInfoList, err = ConsultationReplyList(consultationId)
	if err != nil {
		t.Fatal(err)
	}
	if len(replyInfoList) != 0 {
		t.Fatal(errors.New("Consultation不应存在回复"))
	}
}

func testSetStatus(consultationId int, t *testing.T) {
	consul := &Consultation{Id: consultationId}
	consul.Status = enum.DONE
	consul.Update("status")
	consulWithUserInfo, err := ConsultationGetWithUserInfo(consul.Id)
	if err != nil {
		t.Fatal(err)
	}
	if consulWithUserInfo == nil || consulWithUserInfo.Status != consul.Status {
		t.Fatal("修改Consultation状态失败")
	}
}

func testConsultationGetWithUserInfo(consultationId int, t *testing.T) {
	consulWithUserInfo, err := ConsultationGetWithUserInfo(consultationId)
	if err != nil {
		t.Fatal(err)
	}
	if consulWithUserInfo == nil {
		t.Fatal("查询ConsultationGetWithUser失败。")
	}
}
func testGetUnexixtConsultation(unexixtConsultationId int, t *testing.T) {
	consulPlus, err := ConsultationGetWithUserInfo(unexixtConsultationId)
	if err != nil {
		t.Fatal(err)
	}
	if consulPlus != nil {
		t.Fatal(errors.New("该Consultation不应存在"))
	}
}
