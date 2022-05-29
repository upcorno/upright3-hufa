package dao

import (
	"errors"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//合作意向（维权、监测）用户提交信息
type CooperationIntention struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') CHAR(16)" json:"name"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone"`
	Organization    string    `xorm:"not null comment('组织结构') VARCHAR(60) default('')" json:"organization"`
	Description     string    `xorm:"not null comment('意向描述') TEXT default('')" json:"description"`
	Resume          string    `xorm:"not null comment('权利概要') TEXT default('')" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50) default('')" json:"customer_address"`
	DealRemark      string    `xorm:"comment('回访时备注') TEXT default('')" json:"deal_remark"`
	Category        string    `xorm:"comment('合作类型') unique(category_creator_uid) index CHAR(16)" json:"category"`
	CreatorUid      int       `xorm:"not null unique(category_creator_uid) comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type cooperationIntentionDao struct{}

var CooperationDao *cooperationIntentionDao

func (d *cooperationIntentionDao) Insert(r *CooperationIntention) (id int, err error) {
	if r.Name == "" || r.Category == "" || r.Phone == "" || r.CreatorUid == 0 {
		err = errors.New("必须指定Name、Category、Phone、CreatorUid字段")
		return
	}
	if utf8.RuneCountInString(r.Name) > 16 || utf8.RuneCountInString(r.Phone) > 20 {
		err = errors.New("Name不可超过16个字符、且Phone不超过20个字符")
		return
	}
	if r.CreateTime == 0 {
		r.CreateTime = int(time.Now().Unix())
	}
	_, err = Db.InsertOne(r)
	if err == nil {
		id = r.Id
	}
	return
}

func (d *cooperationIntentionDao) Get(id int, category string, creatorUid int) (has bool, r *CooperationIntention, err error) {
	condiBean := &CooperationIntention{
		Id:         id,
		Category:   category,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && !(condiBean.Category != "" && condiBean.CreatorUid != 0) {
		err = errors.New("需指定CooperationIntention的Id或Category与CreatorUid字段")
		return
	}
	has, err = Db.Get(condiBean)
	if condiBean.CreateTime != 0 {
		r = condiBean
	}
	return
}

func (d *cooperationIntentionDao) Update(id int, category string, creatorUid int, r *CooperationIntention, columns ...string) (err error) {
	condiBean := &CooperationIntention{
		Id:         id,
		Category:   category,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && !(condiBean.Category != "" && condiBean.CreatorUid != 0) {
		err = errors.New("需指定CooperationIntention的Id或Category与CreatorUid字段")
		return
	}
	_, err = Db.Cols(columns...).Update(r, condiBean)
	return
}

func (d *cooperationIntentionDao) delete(id int, category string, creatorUid int) (err error) {
	condiBean := &CooperationIntention{
		Id:         id,
		Category:   category,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && !(condiBean.Category != "" && condiBean.CreatorUid != 0) {
		err = errors.New("需指定CooperationIntention的Id或Category与CreatorUid字段")
		return
	}
	_, err = Db.Delete(condiBean)
	return
}

type CooperationSearchParams struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type searchResult struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

func (d *cooperationIntentionDao) BackendList(category string, page *Page, search *CooperationSearchParams) (result *PageResult, err error) {
	if category == "" {
		err = errors.New("必须指定合作类型")
		return
	}
	searchInfo := []searchResult{}
	sess := Db.NewSession()
	sess.Table("cooperation_intention")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	sess.Where("category = ?", category)
	d.dealSearch(sess, search)
	pageResult, err := page.GetResults(sess, &searchInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func (d *cooperationIntentionDao) dealSearch(sess *xorm.Session, search *CooperationSearchParams) {
	if search.DealResult != "" {
		sess.Where("deal_result = ?", search.DealResult)
	}
	if search.CustomerAddress != "" {
		sess.Where("customer_address like ?", "%"+search.CustomerAddress+"%")
	}
	if search.CreateTimeMin != 0 {
		sess.Where("create_time >= ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("create_time <= ?", search.CreateTimeMax)
	}
}

func (d *cooperationIntentionDao) CountNewItems(minId int, category string) (count int, maxId int, err error) {
	type id struct {
		Id int `json:"id"`
	}
	ids := []id{}
	err = Db.Table("cooperation_intention").
		Cols("id").
		Where("id > ?", minId).
		And("category = ?", category).
		Desc("id").
		Find(&ids)
	if err != nil {
		return
	}
	count = len(ids)
	if count > 0 {
		maxId = ids[0].Id
	} else {
		maxId = minId
	}
	return
}
