package dao

import (
	"errors"
	"testing"
)

var testRightsProtection *RightsProtection

func TestRightsProtection(t *testing.T) {
	defer testDeleteRightsProtection(t)
	testRightsProtection = testAddRightsProtection(t)
	testGetAndUpdateInfo(t)
}

func testAddRightsProtection(t *testing.T) (r *RightsProtection) {
	r = &RightsProtection{
		Name:       "name",
		CreatorUid: TestUserId,
	}
	_, err := RightsProtectionDao.Insert(r)
	if err == nil {
		t.Fatal(errors.New("Name、Phone、CreatorUid三者为必填字段"))
	}
	r.Phone = "ddddd"
	_, err = RightsProtectionDao.Insert(r)
	if err != nil {
		t.Fatal(err)
	}
	_, err = RightsProtectionDao.Insert(r)
	if err == nil {
		t.Fatal(errors.New("一个用户只能填一个维权意向"))
	}
	return
}

func testGetAndUpdateInfo(t *testing.T) {
	has, tmp, err := RightsProtectionDao.Get(testRightsProtection.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询RightsProtection失败"))
	}
	tmp.CustomerAddress = "new ppp"
	tmp.Name = "new name"
	RightsProtectionDao.Update(testRightsProtection.Id, 0, tmp, "name")
	has, newTmp, err := RightsProtectionDao.Get(testRightsProtection.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询RightsProtection失败"))
	}
	if newTmp.CustomerAddress == tmp.CustomerAddress {
		t.Fatal(errors.New("RightsProtection.CustomerAddress不应该被改变"))
	}
	if newTmp.Name != tmp.Name {
		t.Fatal(errors.New("RightsProtection.Name 应该被改变"))
	}
}

func testDeleteRightsProtection(t *testing.T) {
	err := RightsProtectionDao.delete(testRightsProtection.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
}
