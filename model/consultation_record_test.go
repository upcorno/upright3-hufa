package model

import (
	"testing"
	"law/enum"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestConsultationRecordCreate(t *testing.T) {
	record := &ConsultationRecord{
		ConsultationId:  40,
		CommunicatorUid: 1094,
		Type:            enum.QUERY,
		Content:         "测试提问",
		CreateTime:      int(time.Now().Unix()),
	}
	err := ConsultationRecordCreate(record)
	assert.Nil(t, err)
	count, err := Db.Delete(&ConsultationRecord{
		Id: record.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}

func TestConsultationRecordList(t *testing.T) {
	consultationId := 40
	list, err := ConsultationRecordList(consultationId)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(list), 1)
}