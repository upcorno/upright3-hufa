package service

import (
	"fmt"
	"law/model"
	"regexp"

	"xorm.io/xorm"
)

type LegalIssueSearch struct {
	SearchText string `json:"search_text" form:"search_text" query:"search_text"`
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

	legalIssueRes, ok  := pageResult.Rows.(*[]model.LegalIssue)
	if !ok {
		addRes, err := model.LegalIssueListByRand((page.ItemNum))
		if err != nil {
			return nil, err
		}
		return &model.PageResult{Rows: addRes, Total: 0}, nil
	}
	if len(*legalIssueRes) < page.ItemNum {
		addRes, err := model.LegalIssueListByRand((page.ItemNum-len(*legalIssueRes)))
		if err != nil {
			return nil, err
		}
		*legalIssueRes = append(*legalIssueRes, addRes...)
	}
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
