package model 

import (
	_ "github.com/go-sql-driver/mysql"
)

//收藏普法问题
func FavoritesAdd(favorites *Favorites) error {
	_, err := Db.InsertOne(favorites)
	return err
}

//取消收藏普法问题
func FavoritesCancel(uid int, issueId int) error {
	_, err := Db.Delete(&Favorites{
		UserId: uid,
		IssueId: issueId,
	})
	return err
}

//问题是否收藏
func IssueIsFavorites(uid int, issueId int) (bool, error) {
	has, err := Db.Exist(&Favorites{
		UserId: uid,
		IssueId: issueId,
	})
	return has, err
}