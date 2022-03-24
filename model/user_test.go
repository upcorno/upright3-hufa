package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := User{}
	user.AppId = fmt.Sprintf("appid_test(%d)", time.Now().Unix())
	user.Openid = fmt.Sprintf("Openid(%d)", time.Now().Unix())
	assert.Nil(t, user.Insert())
	has, err := user.Get()
	assert.Nil(t, err)
	assert.True(t, has)
	openid := fmt.Sprintf("update(%d)", time.Now().Unix())
	user.Openid = openid
	assert.Nil(t, user.Update())
	user.Get()
	assert.Equal(t, openid, user.Openid)
}
