package service

import (
	"fmt"
	"law/model"
	"regexp"

	"xorm.io/xorm"
)

type LegalIssueSearch struct {
	SearchText         string   `json:"search_text" form:"search_text" query:"search_text"`
}

//法律知识搜索
func LegalIssueList(page *model.Page, search *LegalIssueSearch) (*model.PageResult, error) {
	legalIssues := []model.LegalIssue{}
	sess := model.Db.NewSession()
	sess.Table("legal_issue")
	dealSearch(sess, search)
	pageResult, err := page.GetResults(sess, &legalIssues)
	if err != nil {
		return nil, err
	}
	//todo
	// legalIssueRes := pageResult.Rows.(*[]model.LegalIssue)
	// if len(*legalIssueRes) < page.ItemNum {
	// 	fmt.Println("++++", legalIssueRes)
	// 		智能填充数据

	// }
	return pageResult, err
}

func dealSearch(sess *xorm.Session, search *LegalIssueSearch) {
	if search.SearchText != "" {
		var seachMode string
		if regexp.MustCompile(`(\s[\+,\-,\~,\>,\<])|(^[\+,\-,\~,\>,\<])|(\S\*)`).MatchString(search.SearchText) {
			seachMode = "IN BOOLEAN MODE"
		} else {
			seachMode = "IN NATURAL LANGUAGE MODE"
		}
		sess.Where(fmt.Sprintf("MATCH(title, issue_text) AGAINST (? %s)", seachMode), search.SearchText)
	}
}