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
	_, err := Db.Delete(Favorites{
		UserId: uid,
		IssueId: issueId,
	})
	return err
}