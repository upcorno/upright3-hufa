package model

import (
	"law/enum"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestConsultationCreate(t *testing.T) {
	consultation := &Consultation{
		Question: "单元测试问题",
		ConsultantUid: 1094,
		Status: enum.DOING,
		CreateTime:      int(time.Now().Unix()),
	}
	err := ConsultationCreate(consultation)
	assert.Nil(t, err)
	count, err := Db.Delete(&Consultation{
		Id: consultation.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}

func TestConsultationInfoGet(t *testing.T) {
	consultationId := 41
	consultationInfo, err := ConsultationInfoGet(consultationId)
	assert.Nil(t, err)
	status := consultationInfo.Status
	err = ConsultationStatusSet(consultationId, enum.TOSERVE)
	assert.Nil(t, err)
	consultationInfoUpdate, err := ConsultationInfoGet(consultationId)
	assert.Nil(t, err)
	assert.Equal(t, consultationInfoUpdate.Status, enum.TOSERVE)
	err = ConsultationStatusSet(consultationId, status)
    assert.Nil(t, err)
}

func TestConsultationList(t *testing.T) {
	uid := 1098
	consultationList, err := ConsultationList(uid)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(consultationList), 1)

}