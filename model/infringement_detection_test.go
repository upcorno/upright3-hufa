package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInfringementDetectionAdd(t *testing.T) {
	infringementDetection := &InfringementDetection{
		CreatorUid: 1094,
		Name: "测试",
		Phone: "15634567854",
		Organization: "上海右上角垦丁联合研究院",
		Description: "我的米老鼠描述",
		Resume: "侵权监测概要",
		CreateTime: int(time.Now().Unix()),
	}
	err := InfringementDetectionAdd(infringementDetection)
	assert.Nil(t, err)
	infringementDetectionGet, err := InfringementDetectionGet(infringementDetection.Id)
	assert.Nil(t, err)
	assert.Equal(t, infringementDetection.Name, infringementDetectionGet.Name)
	assert.Equal(t, infringementDetection.Phone, infringementDetectionGet.Phone)
	count, err := Db.Delete(&InfringementDetection{
		Id: infringementDetectionGet.Id,
	})
	assert.Equal(t, int64(1), count)
	assert.Nil(t, err)
}