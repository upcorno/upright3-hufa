package model

import (
	"law/utils"

	_ "github.com/go-sql-driver/mysql"
)

//获取普法问题详情
func LegalIssueGet(legalIssueId int) (LegalIssue, error) {
	issue := LegalIssue{}
	_, err := Db.Table("legal_issue").Where("id=?", legalIssueId).Get(&issue)
	return issue, err
}

//获取该分类查询下面的普法问题
func LegalIssueListByCategory(categoryId int) ([]LegalIssue, error) {
	legalIssueList := []LegalIssue{}
	err := Db.Table("legal_issue").Where("category_id=?", categoryId).Find(&legalIssueList)
	return legalIssueList, err
}

//随机获取普法问题
func LegalIssueListByRand(num int, exceptArr []int) ([]LegalIssue, error) {
	legalIssueList := []LegalIssue{}
	err := Db.Table("legal_issue").In("id", utils.RandSlice(num, exceptArr)).Find(&legalIssueList)
	return legalIssueList, err
}
