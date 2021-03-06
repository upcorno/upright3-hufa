package dao

import (
	"errors"
	"law/enum"
	"testing"
)

var testCooperationIntention *CooperationIntention

func TestCooperationIntention(t *testing.T) {
	defer testDeleteCooperationIntention(t)
	testCooperationIntention = testAddCooperationIntention(t)
	testGetAndUpdate(t)
}

func testAddCooperationIntention(t *testing.T) (r *CooperationIntention) {
	category := enum.PROTECT
	r = &CooperationIntention{
		Name:       "name",
		CreatorUid: TestUserId,
	}
	_, err := CooperationDao.Insert(category, r)
	if err == nil {
		t.Fatal(errors.New("Name、Phone、CreatorUid三者为必填字段"))
	}
	r.Phone = "d333ddss66dd11g2122号"
	_, err = CooperationDao.Insert(category, r)
	if err != nil {
		t.Fatal(err)
	}
	_, err = CooperationDao.Insert(category, r)
	if err == nil {
		t.Fatal(errors.New("一个用户只能填一个维权意向"))
	}
	return
}

func testGetAndUpdate(t *testing.T) {
	has, tmp, err := CooperationDao.Get(enum.PROTECT, testCooperationIntention.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询CooperationIntention失败"))
	}
	tmp.CustomerAddress = "new ppp"
	tmp.Name = "new name"
	CooperationDao.Update(enum.PROTECT, testCooperationIntention.Id, 0, tmp, "name")
	has, newTmp, err := CooperationDao.Get(enum.PROTECT, testCooperationIntention.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("查询CooperationIntention失败"))
	}
	if newTmp.CustomerAddress == tmp.CustomerAddress {
		t.Fatal(errors.New("CooperationIntention.CustomerAddress不应该被改变"))
	}
	if newTmp.Name != tmp.Name {
		t.Fatal(errors.New("CooperationIntention.Name 应该被改变"))
	}
}

func testDeleteCooperationIntention(t *testing.T) {
	err := CooperationDao.delete(enum.PROTECT, testCooperationIntention.Id, 0)
	if err != nil {
		t.Fatal(err)
	}
}
