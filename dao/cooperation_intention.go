package dao

import (
	"errors"
	"law/enum"
	"time"
	"unsafe"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//合作意向（维权、监测）用户提交信息
type CooperationIntention struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Name            string    `xorm:"not null comment('姓名') CHAR(16)" json:"name" validate:"required,min=1,max=16"`
	Phone           string    `xorm:"not null comment('电话号码') CHAR(20)" json:"phone" validate:"required,min=1,max=20"`
	Organization    string    `xorm:"not null comment('组织结构') VARCHAR(60) default('')" json:"organization"`
	Description     string    `xorm:"not null comment('意向描述') TEXT default('')" json:"description"`
	Resume          string    `xorm:"not null comment('权利概要') TEXT default('')" json:"resume"`
	DealResult      string    `xorm:"not null comment('处理状态:未回访 有合作意向 无合作意向 已合作') VARCHAR(10) default('未回访')" json:"deal_result"`
	CustomerAddress string    `xorm:"not null comment('回访时记录客户地址') VARCHAR(50) default('')" json:"customer_address"`
	DealRemark      string    `xorm:"not null comment('回访时备注') TEXT default('')" json:"deal_remark"`
	CreatorUid      int       `xorm:"not null unique comment('创建人id') index UNSIGNED INT" json:"creator_uid" validate:"required,min=1"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}
type rightsProtection struct {
	CooperationIntention `xorm:"extends"`
}
type infringementMonitor struct {
	CooperationIntention `xorm:"extends"`
}

var cooperationTableMap map[enum.Cooperation]string = map[enum.Cooperation]string{enum.PROTECT: "rights_protection", enum.MONITOR: "infringement_monitor"}

type cooperationIntentionDao struct{}

var CooperationDao *cooperationIntentionDao

func (d *cooperationIntentionDao) convertToTarget(category enum.Cooperation, r *CooperationIntention) (bean interface{}, err error) {
	switch category {
	case enum.PROTECT:
		bean = (*rightsProtection)(unsafe.Pointer(r))
	case enum.MONITOR:
		bean = (*infringementMonitor)(unsafe.Pointer(r))
	default:
		err = errors.New("不支持的合作类型")
		return
	}
	return
}

func (d *cooperationIntentionDao) Insert(category enum.Cooperation, r *CooperationIntention) (id int, err error) {
	err = validator.New().Struct(r)
	if err != nil {
		return
	}
	if r.CreateTime == 0 {
		r.CreateTime = int(time.Now().Unix())
	}
	bean, err := d.convertToTarget(category, r)
	if err != nil {
		return
	}
	_, err = Db.InsertOne(bean)
	if err == nil {
		id = r.Id
	}
	return
}

func (d *cooperationIntentionDao) Get(category enum.Cooperation, id int, creatorUid int) (has bool, r *CooperationIntention, err error) {
	tmp := &CooperationIntention{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if tmp.Id == 0 && tmp.CreatorUid == 0 {
		err = errors.New("需指定Id或CreatorUid字段")
		return
	}
	condiBean, err := d.convertToTarget(category, tmp)
	if err != nil {
		return
	}
	has, err = Db.Get(condiBean)
	if tmp.CreateTime != 0 {
		r = tmp
	}
	return
}

func (d *cooperationIntentionDao) Update(category enum.Cooperation, id int, creatorUid int, r *CooperationIntention, columns ...string) (err error) {
	tmp := &CooperationIntention{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if tmp.Id == 0 && tmp.CreatorUid == 0 {
		err = errors.New("需指定Id或CreatorUid字段")
		return
	}
	condiBean, err := d.convertToTarget(category, tmp)
	if err != nil {
		return
	}
	updateBean, err := d.convertToTarget(category, r)
	if err != nil {
		return
	}
	_, err = Db.Cols(columns...).Update(updateBean, condiBean)
	return
}

func (d *cooperationIntentionDao) delete(category enum.Cooperation, id int, creatorUid int) (err error) {
	tmp := &CooperationIntention{
		Id:         id,
		CreatorUid: creatorUid,
	}
	if tmp.Id == 0 && tmp.CreatorUid == 0 {
		err = errors.New("需指定Id或CreatorUid字段")
		return
	}
	condiBean, err := d.convertToTarget(category, tmp)
	if err != nil {
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

func (d *cooperationIntentionDao) BackendList(category enum.Cooperation, page *Page, search *CooperationSearchParams) (result *PageResult, err error) {
	table, ok := cooperationTableMap[category]
	if !ok {
		err = errors.New("不支持的合作类型")
		return
	}
	searchInfo := []searchResult{}
	sess := Db.NewSession()
	sess.Table(table)
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

func (d *cooperationIntentionDao) CountNewItems(minId int, category enum.Cooperation) (count int, maxId int, err error) {
	table, ok := cooperationTableMap[category]
	if !ok {
		err = errors.New("不支持的合作类型")
		return
	}
	type id struct {
		Id int `json:"id"`
	}
	ids := []id{}
	err = Db.Table(table).
		Cols("id").
		Where("id > ?", minId).
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
