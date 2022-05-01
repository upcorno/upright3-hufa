package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//收藏普法问题
func LegalIssueAddFavorite(uid int, issueId int) error {
	favorite := &LegalIssueFavorite{}
	favorite.IssueId = issueId
	favorite.UserId = uid
	has, err := Db.Exist(&LegalIssueFavorite{
		UserId:  favorite.UserId,
		IssueId: favorite.IssueId,
	})
	if err != nil {
		return err
	}
	if has {
		return nil
	}
	_, err = Db.InsertOne(favorite)
	return err
}

//取消收藏普法问题
func LegalIssueCancelFavorite(uid int, issueId int) error {
	_, err := Db.Delete(&LegalIssueFavorite{
		UserId:  uid,
		IssueId: issueId,
	})
	return err
}

//问题是否收藏
func LegalIssueIsFavorite(uid int, issueId int) (bool, error) {
	has, err := Db.Exist(&LegalIssueFavorite{
		UserId:  uid,
		IssueId: issueId,
	})
	return has, err
}
