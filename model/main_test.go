package model

import (
	"law/conf"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	conf.Init()
	Init()
	code := m.Run()
	os.Exit(code)
}
