package model

import (
	"law/enum"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDetectionRetureVisitAdd(t *testing.T) {
	detectionRetureVisit := &DetectionReturnVisit{
		CreatorUid: 1094,
		DetectionId: 4,
		Classification: enum.NORETURN,
		CreateTime: int(time.Now().Unix()),
	}
	err := DetectionRetureVisitAdd(detectionRetureVisit)
	assert.Nil(t, err)
	detectionRetureVisitGet, err := DetectionReturnVisitGet(detectionRetureVisit.DetectionId)
	assert.IsType(t,  DetectionReturnVisit{} ,detectionRetureVisitGet)
    assert.Equal(t, detectionRetureVisit.Classification, detectionRetureVisitGet.Classification)
	assert.Nil(t, err)
	detectionRetureVisitUpdate := &DetectionReturnVisit{
		DetectionId: detectionRetureVisit.DetectionId,
		Classification: enum.HAVEINTENTION,
		CustomerAddress: "安徽省安庆市怀宁县",
		Remark: "张冠军备注",
	}
	err = DetectionRetureVisitUpdate(detectionRetureVisitUpdate)
	assert.Nil(t, err)
	detectionRetureVisitGet, err = DetectionReturnVisitGet(detectionRetureVisit.DetectionId)
	assert.Nil(t, err)
	assert.Equal(t, detectionRetureVisitUpdate.Classification, detectionRetureVisitGet.Classification)
	count, err := Db.Delete(&DetectionReturnVisit{
		Id: detectionRetureVisit.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}

