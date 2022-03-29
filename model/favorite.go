package model

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

//收藏普法问题
func FavoriteAdd(favorite *Favorite) error {
	has, err := Db.Exist(&Favorite{
		UserId:  favorite.UserId,
		IssueId: favorite.IssueId,
	})
	if err != nil {
		return err
	}
	if has {
		return errors.New("该普法问题已被收藏，请勿重复收藏！")
	} else {
		_, err = Db.InsertOne(favorite)
		return err
	}
}

//取消收藏普法问题
func FavoriteCancel(uid int, issueId int) error {
	_, err := Db.Delete(&Favorite{
		UserId:  uid,
		IssueId: issueId,
	})
	return err
}

//问题是否收藏
func IssueIsFavorite(uid int, issueId int) (bool, error) {
	has, err := Db.Exist(&Favorite{
		UserId:  uid,
		IssueId: issueId,
	})
	return has, err
}

//用户收藏列表
func FavoriteList(uid int) ([]LegalIssue, error) {
	legalIssueList := []LegalIssue{}
	err := Db.Table("legal_issue").
		Join("INNER", "favorite", "favorite.issue_id = legal_issue.id").
		Where("favorite.user_id = ?", uid).
		Find(&legalIssueList)
	return legalIssueList, err
}
