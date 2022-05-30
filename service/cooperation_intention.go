package service

import (
	"errors"
	dao "law/dao"
	"law/enum"
	"time"
)

type cooperationSrv struct{}

var CooperationSrv = &cooperationSrv{}

type CooperationDealInfo struct {
	Id              int    `json:"id" query:"id" form:"id" validate:"required,gt=0"`
	DealResult      string `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string `json:"customer_address" form:"customer_address" query:"customer_address" validate:"min=0,max=50"`
	DealRemark      string `json:"deal_remark" form:"deal_remark" query:"deal_remark" validate:"min=0"`
}

func (p *cooperationSrv) SetDealInfo(category enum.Cooperation, id int, dealInfo *CooperationDealInfo) (err error) {
	bean := &dao.CooperationIntention{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: dealInfo.CustomerAddress,
		DealRemark:      dealInfo.DealRemark,
	}
	err = dao.CooperationDao.Update(
		category,
		id,
		0,
		bean,
		"deal_result",
		"customer_address",
		"deal_remark",
	)
	return
}

type CooperationBaseInfo struct {
	Name         string `json:"name" form:"name" query:"name" validate:"min=1,max=16"`
	Phone        string `json:"phone" form:"phone" query:"phone" validate:"min=3,max=20"`
	Organization string `json:"organization" form:"organization" query:"organization" validate:"min=0,max=60"`
	Description  string `json:"description" form:"description" query:"description" validate:"min=0"`
	Resume       string `json:"resume" form:"resume" query:"resume" validate:"min=0"`
}

func (p *cooperationSrv) UpdateBaseInfo(category enum.Cooperation, creatorUid int, baseInfo *CooperationBaseInfo) (err error) {
	bean := &dao.CooperationIntention{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
	}
	err = dao.CooperationDao.Update(category, 0, creatorUid, bean, "name", "phone", "organization", "description", "resume")
	return
}

func (p *cooperationSrv) Add(baseInfo *CooperationBaseInfo, category enum.Cooperation, creatorUid int) (id int, err error) {
	has, _, err := dao.CooperationDao.Get(category, 0, creatorUid)
	if err != nil {
		return
	}
	if has {
		err = errors.New("系统已存在记录，请勿重复添加！")
		return
	}
	bean := &dao.CooperationIntention{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
		CreatorUid:   creatorUid,
		CreateTime:   int(time.Now().Unix()),
	}
	id, err = dao.CooperationDao.Insert(category, bean)
	return
}

func (p *cooperationSrv) BgGet(category enum.Cooperation, id int) (bean *dao.CooperationIntention, err error) {
	has, bean, err := dao.CooperationDao.Get(category, id, 0)
	if err != nil {
		return
	}
	if !has {
		err = errors.New("无查询的 CooperationIntention")
		return
	}
	return
}

func (p *cooperationSrv) Get(category enum.Cooperation, creatorUid int) (bean *dao.CooperationIntention, err error) {
	_, bean, err = dao.CooperationDao.Get(category, 0, creatorUid)
	return
}
