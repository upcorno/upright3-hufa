package service

import (
	"fmt"
	"law/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	user := model.User{}
	user.AppId = fmt.Sprintf("appid_test(%d)", time.Now().Unix())
	user.Openid = fmt.Sprintf("Openid(%d)", time.Now().Unix())
	assert.Nil(t, user.Insert())
	SetNameAndAvatarUrl(user.Id, "name", "url")
	has, err := user.Get()
	assert.Nil(t, err)
	assert.True(t, has)
	assert.Equal(t, "name", user.NickName)
	assert.Equal(t, "url", user.AvatarUrl)
}
