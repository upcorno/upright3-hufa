package model

import (
	_ "github.com/go-sql-driver/mysql"
)


//获取普法问题详情
func LegalIssueGet(legalIssueId int) (LegalIssue, error) {
	issue := LegalIssue{}
	_, err := Db.Table("legal_issue").Where("id=?", legalIssueId).Get(&issue)
	return issue, err
}