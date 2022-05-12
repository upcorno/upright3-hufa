package model

import (
	"errors"
	"testing"
)

var testRightsProtection RightsProtection

func TestRightsProtection(t *testing.T) {
	testRightsProtection = testAddRightsProtection(t)
	testGetAndUpdateInfo(t)
	testDeleteRightsProtection(t)
}

func testAddRightsProtection(t *testing.T) (r RightsProtection) {
	r.Name = "name"
	r.CreatorUid = TestUid
	err := r.Insert()
	if err == nil {
		t.Fatal(errors.New("Name、Phone、CreatorUid三者为必填字段"))
	}
	r.Phone = "ddddd"
	err = r.Insert()
	if err != nil {
		t.Fatal(err)
	}
	err = r.Insert()
	if err == nil {
		t.Fatal(errors.New("一个用户只能填一个维权意向"))
	}
	return
}

func testGetAndUpdateInfo(t *testing.T) {
	tmp := &RightsProtection{Id: testRightsProtection.Id}
	has, err := tmp.Get()
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询RightsProtection失败"))
	}
	tmp.CustomerAddress = "new ppp"
	tmp.Name = "new name"
	tmp.Update("name")
	newTmp := &RightsProtection{Id: testRightsProtection.Id}
	has, err = newTmp.Get()
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询RightsProtection失败"))
	}
	if newTmp.CustomerAddress == tmp.CustomerAddress {
		t.Fatal(errors.New("RightsProtection.CustomerAddress不应该被改变"))
	}
	if newTmp.Name == testRightsProtection.Name {
		t.Fatal(errors.New("RightsProtection.Name 应该被改变"))
	}
	testRightsProtection.Update()
}

func testDeleteRightsProtection(t *testing.T) {
	err := testRightsProtection.delete()
	if err != nil {
		t.Fatal(err)
	}
}
