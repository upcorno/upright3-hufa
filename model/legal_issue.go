package model

import (
	"fmt"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//获取普法问题详情
func LegalIssueGet(legalIssueId int) (LegalIssue, error) {
	issue := LegalIssue{}
	_, err := Db.Table("legal_issue").Where("id=?", legalIssueId).Get(&issue)
	return issue, err
}

type LegalIssueSearch struct {
	SearchText     string `json:"search_text" form:"search_text" query:"search_text"`
	FirstCategory  string `json:"first_category" form:"first_category" query:"first_category"`
	SecondCategory string `json:"second_category" form:"second_category" query:"second_category"`
	//FavoriteUid本来是数字，查询时为了方便直接定义为字符串
	FavoriteUid int
	IsFavorite  bool `json:"is_favorite" form:"is_favorite" query:"is_favorite"`
	InSummary   bool `json:"in_summary" form:"in_summary" query:"in_summary"`
}

func LegalIssueList(page *Page, search *LegalIssueSearch) (*PageResult, error) {
	legalIssues := new([]LegalIssue)
	sess := Db.NewSession()
	sess.Table("legal_issue")
	if search.InSummary {
		sess.Cols("legal_issue.id", "creator_uid", "first_category", "second_category", "tags", "title")
	} else {
		sess.Cols("legal_issue.id", "creator_uid", "first_category", "second_category", "tags", "title", "imgs", "content")
	}
	dealSearch(sess, search)
	pageResult, err := page.GetResults(sess, legalIssues)
	if err != nil {
		return nil, err
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
		sess.Where(fmt.Sprintf("MATCH(search_text) AGAINST (? %s)", seachMode), search.SearchText)
	}
	if search.FirstCategory != "" {
		sess.Where("first_category = ?", search.FirstCategory)
	}
	if search.SecondCategory != "" {
		sess.Where("second_category = ?", search.SecondCategory)
	}
	if search.IsFavorite {
		sess.
			Join("INNER", "legal_issue_favorite", "legal_issue_favorite.issue_id = legal_issue.id").
			Where("legal_issue_favorite.user_id = ?", search.FavoriteUid)
	}
}

func IssueCategoryList() ([]map[string][]string, error) {
	sql := "SELECT distinct second_category,first_category FROM legal_issue order by first_category desc"
	//希望著作权在前面，所以sql中排了序
	results, err := Db.QueryString(sql)
	categoryList := []map[string][]string{}
	firstCategory := ""
	var category map[string][]string
	for _, item := range results {
		if firstCategory != item["first_category"] {
			firstCategory = item["first_category"]
			category = map[string][]string{}
			category[item["first_category"]] = []string{}
			categoryList = append(categoryList, category)
		}
		category[item["first_category"]] = append(category[item["first_category"]], item["second_category"])
	}
	return categoryList, err
}
