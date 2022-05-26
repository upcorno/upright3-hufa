package dao

import (
	"errors"
	"testing"
)

func TestLigalIssue(t *testing.T) {
	issue := addLegalIssue(t)
	issueId := issue.Id
	newIssue := &LegalIssue{
		Id: issueId,
	}
	has, err := newIssue.Get()
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("legal_issue应该存在"))
	}
	if newIssue.SecondCategory != issue.SecondCategory {
		t.Fatal(errors.New("legal_issue未知错误"))
	}
	testFavorite(issueId, t)
	testLigalIssueList(t)
	testLigalIssueCategoryList(t)
}

var secondCategory string = "二级分类"

func addLegalIssue(t *testing.T) (issue *LegalIssue) {
	issue = &LegalIssue{
		CreatorUid:     TestUid,
		FirstCategory:  "一级分类",
		SecondCategory: secondCategory,
		Tags:           "string",
		Imgs:           "string",
		Title:          "biaoti内容",
		Content:        "内容",
	}
	err := issue.Insert()
	if err != nil {
		t.Fatal(err)
	}
	if issue.Id < 1 {
		t.Fatal(errors.New("issue未添加成功"))
	}
	return
}

func testLigalIssueList(t *testing.T) {
	page := &Page{ItemNum: 5, PageIndex: 1}
	search := &LegalIssueSearch{SecondCategory: secondCategory}
	result, err := LegalIssueList(page, search)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total < 1 {
		t.Fatal(errors.New("至少应存在一个满足条件的LigalIssue"))
	}
	search = &LegalIssueSearch{BusinessCategory: "版权基础"}
	result, err = LegalIssueList(page, search)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total < 1 {
		t.Fatal(errors.New("至少应存在一个满足条件的LigalIssue"))
	}
}

func testLigalIssueCategoryList(t *testing.T) {
	list, err := LegalIssueCategoryList()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) < 1 {
		t.Fatal(errors.New("查询LegalIssue分类失败"))
	}
}
func testFavorite(issueId int, t *testing.T) {
	issueFavorite := &LegalIssueFavorite{
		IssueId: issueId,
		UserId:  TestUid,
	}
	err := issueFavorite.Insert()
	if err != nil {
		t.Fatal(err)
	}
	has, err := issueFavorite.Exist()
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("LegalIssue应该已添加收藏"))
	}
	tmpIssueFavorite := &LegalIssueFavorite{UserId: issueFavorite.UserId, IssueId: issueFavorite.IssueId}
	err = tmpIssueFavorite.Delete()
	if err != nil {
		t.Fatal(err)
	}
	has, err = issueFavorite.Exist()
	if err != nil {
		t.Fatal(err)
	}
	if has {
		t.Fatal(errors.New("LegalIssue应该已删除收藏"))
	}
}
