package dao

import (
	"testing"
)

var TestUserId int

func TestMain(m *testing.M) {
	//setup
	deleteTestUser()
	addTestUser(m)
	m.Run()
}

func addTestUser(m *testing.M) {
	user := &User{
		AppId:  "test_app_id",
		Openid: "test_openid",
	}
	userId, err := UserDao.Insert(user)
	if err != nil {
		panic(err)
	}
	TestUserId = userId
}

func deleteTestUser() {
	UserDao.delete(0, "test_app_id", "test_openid")
}
