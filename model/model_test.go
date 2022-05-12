package model

import (
	"law/conf"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var TestUid int

//model包内单元测试命令：go test -v -args -c ../config_test.toml
func TestMain(m *testing.M) {
	//setup
	setNullLogger()
	conf.Init()
	Init()
	addTestUser(m)

	m.Run()

	deleteTestUser()
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

func setNullLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logWriter, err := os.OpenFile(os.DevNull, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("无法创建日志文件" + os.DevNull)
	}
	log.Logger = log.Output(logWriter)
}
