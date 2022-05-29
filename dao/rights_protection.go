package dao

import (
	"errors"
	"time"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//“我要维权”用户提交信息
//CreatorUid字段具有唯一性约束。即一个用户只能有一个
type RightsProtection struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') VARCHAR(16)" json:"name"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone"`
	Organization    string    `xorm:"not null comment('组织结构') VARCHAR(60) default('')" json:"organization"`
	Description     string    `xorm:"not null comment('维权意向描述') TEXT default('')" json:"description"`
	Resume          string    `xorm:"not null comment('权利概要') TEXT default('')" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50) default('')" json:"customer_address"`
	DealRemark      string    `xorm:"comment('回访时备注') TEXT default('')" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null unique comment('创建人id') index UNSIGNED INT" json:"creator_uid"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type rightsProtectionDao struct{}

var RightsProtectionDao *rightsProtectionDao

func (d *rightsProtectionDao) Insert(r *RightsProtection) (id int, err error) {
	if r.Name == "" || r.Phone == "" || r.CreatorUid == 0 {
		err = errors.New("必须指定Name、Phone、CreatorUid字段")
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

func (d *rightsProtectionDao) Get(id int, creatorUid int) (has bool, r *RightsProtection, err error) {
	condiBean := &RightsProtection{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && condiBean.CreatorUid == 0 {
		err = errors.New("需指定RightsProtection的Id或CreatorUid字段")
		return
	}
	has, err = Db.Get(condiBean)
	if condiBean.CreateTime != 0 {
		r = condiBean
	}
	return
}

func (d *rightsProtectionDao) Update(id int, creatorUid int, r *RightsProtection, columns ...string) (err error) {
	condiBean := &RightsProtection{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && condiBean.CreatorUid == 0 {
		err = errors.New("需指定RightsProtection的Id或CreatorUid字段")
		return
	}
	_, err = Db.Cols(columns...).Update(r, condiBean)
	return
}

func (d *rightsProtectionDao) delete(id int, creatorUid int) (err error) {
	condiBean := &RightsProtection{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if condiBean.Id == 0 && condiBean.CreatorUid == 0 {
		err = errors.New("需指定RightsProtection的Id或CreatorUid字段")
		return
	}
	_, err = Db.Delete(condiBean)
	return
}

type RightsProtectionSearchParams struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type protectionInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

func (d *rightsProtectionDao) BackendList(page *Page, search *RightsProtectionSearchParams) (*PageResult, error) {
	searchInfo := []protectionInfo{}
	sess := Db.NewSession()
	sess.Table("rights_protection")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	d.dealSearch(sess, search)
	pageResult, err := page.GetResults(sess, &searchInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func (d *rightsProtectionDao) dealSearch(sess *xorm.Session, search *RightsProtectionSearchParams) {
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
