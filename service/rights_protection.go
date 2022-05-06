package service

import (
	"errors"
	"law/model"
	"time"

	"xorm.io/xorm"
)

type RightsProtectionDealInfo struct {
	Id              int     `json:"id" query:"id" form:"id" validate:"required,gt=0"`
	DealResult      string  `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress *string `json:"customer_address" form:"customer_address" query:"customer_address" validate:"min=0,max=50"`
	DealRemark      *string `json:"deal_remark" form:"deal_remark" query:"deal_remark" validate:"min=0"`
}

func RightsProtectionSetDealInfo(id int, dealInfo *RightsProtectionDealInfo) error {
	bean := &model.RightsProtection{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: *dealInfo.CustomerAddress,
		DealRemark:      *dealInfo.DealRemark,
	}
	_, err := model.Db.Cols("deal_result", "customer_address", "deal_remark").Update(bean, &model.RightsProtection{Id: id})
	return err
}

type RightsProtectionBaseInfo struct {
	Name         string  `json:"name" form:"name" query:"name" validate:"min=1,max=16"`
	Phone        string  `json:"phone" form:"phone" query:"phone" validate:"min=3,max=20"`
	Organization *string `json:"organization" form:"organization" query:"organization" validate:"min=0,max=60"`
	Description  *string `json:"description" form:"description" query:"description" validate:"min=0"`
	Resume       *string `json:"resume" form:"resume" query:"resume" validate:"min=0"`
}

func RightsProtectionUpdateBaseInfo(uid int, baseInfo *RightsProtectionBaseInfo) error {
	bean := &model.RightsProtection{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: *baseInfo.Organization,
		Description:  *baseInfo.Description,
		Resume:       *baseInfo.Resume,
	}
	_, err := model.Db.Cols("name", "phone", "organization", "description", "resume").Update(bean, &model.RightsProtection{CreatorUid: uid})
	return err
}

func RightsProtectionAdd(baseInfo *RightsProtectionBaseInfo, creatorUid int) (id int, err error) {
	has, err := model.Db.Exist(&model.RightsProtection{
		CreatorUid: creatorUid,
	})
	if err != nil {
		return
	}
	if has {
		err = errors.New("系统已添加您的侵权监测，请勿重复添加！")
		return
	}
	bean := model.RightsProtection{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: *baseInfo.Organization,
		Description:  *baseInfo.Description,
		Resume:       *baseInfo.Resume,
		CreatorUid:   creatorUid,
		CreateTime:   int(time.Now().Unix()),
	}
	bId, err := model.Db.InsertOne(bean)
	id = int(bId)
	return
}

func RightsProtectionBgGet(beanId int) (model.RightsProtection, error) {
	bean := model.RightsProtection{}
	_, err := model.Db.Table("rights_protection").Where("id=?", beanId).Get(&bean)
	return bean, err
}

func RightsProtectionGet(creatorUid int) (model.RightsProtection, error) {
	bean := model.RightsProtection{}
	_, err := model.Db.Table("rights_protection").Where("creator_uid=?", creatorUid).Get(&bean)
	return bean, err
}

type RightsProtectionSearch struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type rightsProtectionInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

func RightsProtectionBackendList(page *model.Page, search *RightsProtectionSearch) (*model.PageResult, error) {
	searchInfo := []rightsProtectionInfo{}
	sess := model.Db.NewSession()
	sess.Table("rights_protection")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	dealRightsProtectionSearch(sess, search)
	pageResult, err := page.GetResults(sess, &searchInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealRightsProtectionSearch(sess *xorm.Session, search *RightsProtectionSearch) {
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
