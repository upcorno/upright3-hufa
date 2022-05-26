package dao

import (
	"testing"
	"time"
)

var TestUid int

func TestMain(m *testing.M) {
	//setup
	addTestUser(m)
	defer deleteTestUser()
	m.Run()
}

func addTestUser(m *testing.M) {
	user := &User{
		AppId:      "test_app_id",
		Openid:     "test_openid",
		CreateTime: int(time.Now().Unix()),
	}
	err := user.Insert()
	if err != nil {
		panic(err)
	}
	TestUid = user.Id
}

func deleteTestUser() {
	user := &User{
		AppId:  "test_app_id",
		Openid: "test_openid",
	}
	Db.Delete(user)
}
