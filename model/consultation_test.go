package model

import (
	"law/enum"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConsultationCreate(t *testing.T) {
	consultation := &Consultation{
		Question:      "单元测试问题",
		ConsultantUid: 1094,
		Status:        enum.DOING,
		CreateTime:    int(time.Now().Unix()),
	}
	err := ConsultationCreate(consultation)
	assert.Nil(t, err)
	count, err := Db.Delete(&Consultation{
		Id: consultation.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}

func TestConsultationList(t *testing.T) {
	uid := 1098
	consultationList, err := ConsultationList(uid)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(consultationList), 1)
}

func TestConsultationReply(t *testing.T) {
	record := &ConsultationReply{
		ConsultationId:  40,
		CommunicatorUid: 1094,
		Type:            enum.QUERY,
		Content:         "测试提问",
		CreateTime:      int(time.Now().Unix()),
	}
	err := ConsultationAddReply(record)
	assert.Nil(t, err)
	consultationId := 40
	list, err := ConsultationListReply(consultationId)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
	count, err := Db.Delete(&ConsultationReply{
		Id: record.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}
