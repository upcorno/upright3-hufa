package service

import (
	"errors"
	dao "law/dao"
	"time"
)

type protectionSrv struct{}

var Protection = &protectionSrv{}

type RightsProtectionDealInfo struct {
	Id              int    `json:"id" query:"id" form:"id" validate:"required,gt=0"`
	DealResult      string `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string `json:"customer_address" form:"customer_address" query:"customer_address" validate:"min=0,max=50"`
	DealRemark      string `json:"deal_remark" form:"deal_remark" query:"deal_remark" validate:"min=0"`
}

func (p *protectionSrv) SetDealInfo(id int, dealInfo *RightsProtectionDealInfo) (err error) {
	bean := &dao.RightsProtection{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: dealInfo.CustomerAddress,
		DealRemark:      dealInfo.DealRemark,
	}
	err = dao.RightsProtectionDao.Update(
		id,
		0,
		bean,
		"deal_result",
		"customer_address",
		"deal_remark",
	)
	return
}

type RightsProtectionBaseInfo struct {
	Name         string `json:"name" form:"name" query:"name" validate:"min=1,max=16"`
	Phone        string `json:"phone" form:"phone" query:"phone" validate:"min=3,max=20"`
	Organization string `json:"organization" form:"organization" query:"organization" validate:"min=0,max=60"`
	Description  string `json:"description" form:"description" query:"description" validate:"min=0"`
	Resume       string `json:"resume" form:"resume" query:"resume" validate:"min=0"`
}

func (p *protectionSrv) UpdateBaseInfo(creatorUid int, baseInfo *RightsProtectionBaseInfo) (err error) {
	bean := &dao.RightsProtection{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
	}
	err = dao.RightsProtectionDao.Update(0, creatorUid, bean, "name", "phone", "organization", "description", "resume")
	return
}

func (p *protectionSrv) Add(baseInfo *RightsProtectionBaseInfo, creatorUid int) (id int, err error) {
	has, bean, err := dao.RightsProtectionDao.Get(0, creatorUid)
	if err != nil {
		return
	}
	if has {
		err = errors.New("系统已存在记录，请勿重复添加！")
		return
	}
	bean = &dao.RightsProtection{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
		CreatorUid:   creatorUid,
		CreateTime:   int(time.Now().Unix()),
	}
	id, err = dao.RightsProtectionDao.Insert(bean)
	return
}

func (p *protectionSrv) BgGet(id int) (bean *dao.RightsProtection, err error) {
	has, bean, err := dao.RightsProtectionDao.Get(id, 0)
	if err != nil {
		return
	}
	if !has {
		err = errors.New("无查询的RightsProtection")
		return
	}
	return
}

func (p *protectionSrv) Get(creatorUid int) (bean *dao.RightsProtection, err error) {
	_, bean, err = dao.RightsProtectionDao.Get(0, creatorUid)
	return
}
