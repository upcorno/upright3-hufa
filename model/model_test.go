package model

import (
	"testing"
	"time"
)

var TestUid int

//model包内单元测试命令：go test -v -args -c ../config_test.toml
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
	user.Insert()
	TestUid = user.Id
}

func deleteTestUser() {
	user := &User{Id: TestUid}
	Db.Delete(user)
}
