package dao

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//用户收藏
type LegalIssueFavorite struct {
	Id         int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	UserId     int       `xorm:"not null comment('用户id') index UNSIGNED INT" json:"user_id"`
	IssueId    int       `xorm:"not null comment('普法知识问题id') index UNSIGNED INT" json:"issue_id"`
	CreateTime int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

func (f *LegalIssueFavorite) Insert() (err error) {
	if f.IssueId == 0 || f.UserId == 0 {
		err = errors.New("dao:IssueId、UserId不可为空")
		return
	}
	f.CreateTime = int(time.Now().Unix())
	_, err = Db.InsertOne(f)
	return
}

func (f *LegalIssueFavorite) Exist() (has bool, err error) {
	if f.Id == 0 {
		if f.IssueId == 0 || f.UserId == 0 {
			err = errors.New("dao:Id为空时，IssueId、UserId不能为空")
			return
		}
	}
	has, err = Db.Exist(f)
	return
}

func (f *LegalIssueFavorite) Delete() (err error) {
	if f.Id == 0 {
		if !(f.IssueId != 0 && f.UserId != 0) {
			err = errors.New("dao:必须指定Id或同时指定IssueId、UserId。")
			return
		}
	}
	_, err = Db.Delete(&LegalIssueFavorite{
		Id:      f.Id,
		IssueId: f.IssueId,
		UserId:  f.UserId,
	})
	return
}
