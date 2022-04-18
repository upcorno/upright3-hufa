package model

import (
	"law/enum"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProtectionRetureVisitAdd(t *testing.T) {
	protectionRetureVisit := &ProtectionReturnVisit{
		CreatorUid: 1094,
		ProtectionId: 4,
		Classification: enum.NORETURN,
		CreateTime: int(time.Now().Unix()),
	}
	err := ProtectionRetureVisitAdd(protectionRetureVisit)
	assert.Nil(t, err)
	protectionRetureVisitGet, err := ProtectionReturnVisitGet(protectionRetureVisit.ProtectionId)
	assert.IsType(t,  ProtectionReturnVisit{} ,protectionRetureVisitGet)
    assert.Equal(t, protectionRetureVisit.Classification, protectionRetureVisitGet.Classification)
	assert.Nil(t, err)
	protectionRetureVisitUpdate := &ProtectionReturnVisit{
		ProtectionId: protectionRetureVisit.ProtectionId,
		Classification: enum.HAVEINTENTION,
		CustomerAddress: "安徽省安庆市怀宁县",
		Remark: "张冠军备注",
	}
	err = ProtectionRetureVisitUpdate(protectionRetureVisitUpdate)
	assert.Nil(t, err)
	protectionRetureVisitGet, err = ProtectionReturnVisitGet(protectionRetureVisit.ProtectionId)
	assert.Nil(t, err)
	assert.Equal(t, protectionRetureVisitUpdate.Classification, protectionRetureVisitGet.Classification)
	count, err := Db.Delete(&ProtectionReturnVisit{
		Id: protectionRetureVisit.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}