package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//查询当前分类子级分类
func CategoryList(categoryId int) ([]Category, error) {
	// categoryList := []Category{}
	// err := Db.Table("category").Where("pre_category_id=?", categoryId).Find(&categoryList)

	// where 条件中直接用字符串的话，orm 框架就不能发挥到更，可考虑用如下方法代替（暂未测试，仅做示意）
	// orm 框架是业务代码和数据库的中间层，代码中如果大量耦合数据库的字段，不利于维护
	categoryList := []Category{}
	condiBean := Category{
		PreCategoryId: categoryId,
	}
	err := Db.Find(&categoryList, condiBean)
	return categoryList, err
	//不必要的空格可以删掉
}
