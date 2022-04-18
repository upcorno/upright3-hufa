package model

import (
	"law/conf"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	c := make(chan os.Signal, 1)
	conf.Init(c)
	Init()
	code := m.Run()
	os.Exit(code)
}