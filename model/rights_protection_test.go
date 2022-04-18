package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRightsProtectionAdd(t *testing.T) {
	rightsProtection := &RightsProtection{
		CreatorUid: 1094,
		Name: "测试",
		Phone: "15634567854",
		Organization: "上海右上角垦丁联合研究院",
		Description: "我的米老鼠描述",
		Resume: "维权意向概要",
		CreateTime: int(time.Now().Unix()),
	}
	err := RightsProtectionAdd(rightsProtection)
	assert.Nil(t, err)
	rightsProtectionGet, err := RightsProtectionGet(rightsProtection.Id)
	assert.Nil(t, err)
	assert.Equal(t, rightsProtection.Name, rightsProtectionGet.Name)
	assert.Equal(t, rightsProtection.Phone, rightsProtectionGet.Phone)
	count, err := Db.Delete(&RightsProtection{
		Id: rightsProtectionGet.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}