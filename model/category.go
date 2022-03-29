package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//查询当前分类子级分类
func CategoryList(categoryId int) ([]Category, error) {
	categoryList := []Category{}
	err := Db.Table("category").Where("pre_category_id=?", categoryId).Find(&categoryList)
	return categoryList, err

}
