package dao

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 问题“咨询”
type Consultation struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question"`
	Imgs          string    `xorm:"not null comment('描述图片') TEXT  default('')" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT" json:"consultant_uid" validate:"required"`
	Status        string    `xorm:"not null default '处理中' comment('处理中、待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status"`
	CreateTime    int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

type consulDao struct{}

var ConsulDao *consulDao

//创建咨询
func (c *consulDao) Insert(question string, imgs string, consultantUid int, status string) (consulId int, err error) {
	consul := &Consultation{
		Question:      question,
		Imgs:          imgs,
		ConsultantUid: consultantUid,
		Status:        status,
		CreateTime:    int(time.Now().Unix()),
	}
	if consul.Question == "" || consul.ConsultantUid == 0 || consul.Status == "" {
		err = errors.New("Question、ConsultantUid、Status不可以为空值")
		return
	}
	consul.CreateTime = int(time.Now().Unix())
	_, err = Db.InsertOne(consul)
	if err == nil {
		consulId = consul.Id
	}
	return
}

//删除咨询
func (c *consulDao) delete(consulId int) (err error) {
	if consulId == 0 {
		err = errors.New("consulId不可为0")
		return
	}
	_, err = Db.Delete(&Consultation{Id: consulId})
	return
}

func (c *consulDao) Update(consulId int, consul *Consultation, columns ...string) (err error) {
	if consulId == 0 {
		err = errors.New("consulId不可为0")
		return
	}
	_, err = Db.Cols(columns...).Update(consul, &Consultation{Id: consulId})
	return
}

//用户历史咨询记录列表
func (c *consulDao) List(uid int) (consultationList []Consultation, err error) {
	consultationList = []Consultation{}
	err = Db.Table("consultation").
		Where("consultation.consultant_uid = ?", uid).
		Desc("create_time").
		Find(&consultationList)
	return
}

type consultationWithUserInfo struct {
	Id            int    `json:"id"`
	Question      string `json:"question"`
	Imgs          string `json:"imgs"`
	Status        string `json:"status"`
	CreateTime    int    `json:"create_time"`
	ConsultantUid int    `json:"consultant_uid"`
	NickName      string `json:"nick_name"`
	AvatarUrl     string `json:"avatar_url"`
	Phone         string `json:"phone"`
}

//获取咨询信息
func (c *consulDao) GetWithUserInfo(consultationId int) (*consultationWithUserInfo, error) {
	consultationInfo := &consultationWithUserInfo{}
	_, err := Db.Table("consultation").
		Join("INNER", "user", "user.id = consultation.consultant_uid").
		Where("consultation.id=?", consultationId).
		Cols(
			"consultation.id",
			"consultation.question",
			"consultation.imgs",
			"consultation.status",
			"consultation.create_time",
			"consultation.consultant_uid",
			"user.nick_name",
			"user.avatar_url",
			"user.phone",
		).
		Get(consultationInfo)
	if consultationInfo.Id == 0 {
		consultationInfo = nil
	}
	return consultationInfo, err
}

type ConsultationSearchParams struct {
	Status        string `json:"status" query:"status"`
	CreateTimeMin int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax int    `json:"create_time_max" query:"create_time_max"`
}

func (c *consulDao) BackendList(page *Page, search *ConsultationSearchParams) (pageResult *PageResult, err error) {
	type listInfo struct {
		Id       int    `json:"id"`
		Question string `json:"question"`
		NickName string `json:"nick_name"`
		Phone    string `json:"phone"`
		Status   string `json:"status"`
	}
	consultationInfoList := []listInfo{}
	sess := Db.NewSession()
	sess.Table("consultation")
	sess.Join("INNER", "user", "user.id = consultation.consultant_uid")
	sess.Cols(
		"consultation.id",
		"consultation.question",
		"user.nick_name",
		"user.phone",
		"consultation.status",
	)
	c.dealSearch(sess, search)
	pageResult, err = page.GetResults(sess, &consultationInfoList)
	return
}

func (c *consulDao) dealSearch(sess *xorm.Session, search *ConsultationSearchParams) {
	if search.Status != "" {
		sess.Where("consultation.status = ?", search.Status)
	}
	if search.CreateTimeMin != 0 {
		sess.Where("consultation.create_time > ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("consultation.create_time < ?", search.CreateTimeMax)
	}
}
func (c *consulDao) CountNewItems(minId int) (count int, maxId int, err error) {
	type id struct {
		Id int `json:"id"`
	}
	ids := []id{}
	err = Db.Table("consultation").
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
