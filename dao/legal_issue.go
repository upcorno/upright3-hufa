package dao

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//常见知产问题
type LegalIssue struct {
	Id               int       `xorm:"not null pk autoincr INT" json:"id"`
	CreatorUid       int       `xorm:"not null comment('问题创建人id') index UNSIGNED INT" json:"creator_uid"`
	FirstCategory    string    `xorm:"not null comment('一级类别') index CHAR(6)" json:"first_category" validate:"required,min=1,max=6"`
	SecondCategory   string    `xorm:"not null comment('二级类别') index CHAR(25)" json:"second_category" validate:"required,min=1,max=25"`
	BusinessCategory string    `xorm:"not null comment('业务类别') index CHAR(60)" json:"business_category" validate:"required,min=1,max=60"`
	Tags             string    `xorm:"not null comment('问题标签') index VARCHAR(255) default('')" json:"tags" validate:"max=255"`
	Title            string    `xorm:"not null comment('标题') VARCHAR(60)" json:"title" validate:"required,min=1,max=60"`
	Imgs             string    `xorm:"not null comment('普法问题关联图片') TEXT default('')" json:"imgs"`
	Content          string    `xorm:"not null comment('内容') LONGTEXT" json:"content" validate:"required,min=1"`
	SearchText       string    `xorm:"not null comment('全文检索字段') LONGTEXT default('')" json:"-"`
	CreateTime       int       `xorm:"not null UNSIGNED INT default(1651383059)" json:"create_time"`
	UpdateTime       time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}
type legalIssueDao struct{}

var LegalIssueDao *legalIssueDao

func (l *legalIssueDao) Insert(issue *LegalIssue) (issueId int, err error) {
	err = validator.New().Struct(issue)
	if err != nil {
		return
	}
	if issue.CreateTime == 0 {
		issue.CreateTime = int(time.Now().Unix())
	}
	_, err = Db.InsertOne(issue)
	if err == nil {
		issueId = issue.Id
	}
	return
}

func (l *legalIssueDao) Update(issueId int, issue *LegalIssue, columns ...string) (err error) {
	if issueId == 0 {
		err = errors.New("必须指定id。")
		return
	}
	_, err = Db.Cols(columns...).Update(issue, &LegalIssue{Id: issueId})
	return
}

func (l *legalIssueDao) Delete(issueId int) (err error) {
	if issueId == 0 {
		err = errors.New("必须指定id。")
		return
	}
	_, err = Db.Delete(&LegalIssue{Id: issueId})
	return
}

func (l *legalIssueDao) Get(issueId int) (has bool, issue *LegalIssue, err error) {
	issue = &LegalIssue{Id: issueId}
	if issue.Id < 1 {
		err = errors.New("必须指定issueId")
		return
	}
	has, err = Db.Get(issue)
	return
}

func (l *legalIssueDao) MustGet(issueId int) (issue *LegalIssue, err error) {
	has, issue, err := l.Get(issueId)
	if err == nil && !has {
		err = errors.New("此legalIssue不存在。")
	}
	return
}

type LegalIssueSearch struct {
	SearchText       string `json:"search_text" form:"search_text" query:"search_text"`
	FirstCategory    string `json:"first_category" form:"first_category" query:"first_category"`
	SecondCategory   string `json:"second_category" form:"second_category" query:"second_category"`
	BusinessCategory string `json:"business_category" form:"business_category" query:"business_category"`
	FavoriteUid      int
	//deprecated 小程序正式发布后可删除此字段
	IsFavorite   bool `json:"is_favorite" form:"is_favorite" query:"is_favorite"`
	OnlyFavorite bool `json:"only_favorite" form:"only_favorite" query:"only_favorite"`
	InSummary    bool `json:"in_summary" form:"in_summary" query:"in_summary"`
}

func (l *legalIssueDao) List(page *Page, search *LegalIssueSearch) (pageResult *PageResult, err error) {
	legalIssues := new([]LegalIssue)
	sess := Db.NewSession()
	sess.Table("legal_issue")
	cols := []string{"legal_issue.id", "creator_uid", "first_category", "second_category", "business_category", "tags", "title"}
	if !search.InSummary {
		cols = append(cols, "imgs", "content")
	}
	sess.Cols(cols...)
	dealSearch(sess, search)
	pageResult, err = page.GetResults(sess, legalIssues)
	return
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
	if search.BusinessCategory != "" {
		sess.Where("business_category like ?", fmt.Sprintf("%%%s%%", search.BusinessCategory))
	}
	if search.OnlyFavorite {
		sess.
			Join("INNER", "legal_issue_favorite", "legal_issue_favorite.issue_id = legal_issue.id").
			Where("legal_issue_favorite.user_id = ?", search.FavoriteUid)
	}
}

func (l *legalIssueDao) CategoryList() ([]map[string][]string, error) {
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
