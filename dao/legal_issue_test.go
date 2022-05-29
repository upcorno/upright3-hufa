package dao

import (
	"errors"
	"testing"
)

func TestLigalIssue(t *testing.T) {
	issue := addLegalIssue(t)
	issueId := issue.Id
	has, newIssue, err := LegalIssueDao.Get(issueId)
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
		CreatorUid:     TestUserId,
		FirstCategory:  "一级分类",
		SecondCategory: secondCategory,
		Tags:           "string",
		Imgs:           "string",
		Title:          "biaoti内容",
		Content:        "内容",
	}
	issueId, err := LegalIssueDao.Insert(issue)
	if err != nil {
		t.Fatal(err)
	}
	if issueId < 1 {
		t.Fatal(errors.New("issue未添加成功"))
	}
	return
}

func testLigalIssueList(t *testing.T) {
	page := &Page{ItemNum: 5, PageIndex: 1}
	search := &LegalIssueSearch{SecondCategory: secondCategory}
	result, err := LegalIssueDao.List(page, search)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total < 1 {
		t.Fatal(errors.New("至少应存在一个满足条件的LigalIssue"))
	}
	search = &LegalIssueSearch{BusinessCategory: "版权基础"}
	result, err = LegalIssueDao.List(page, search)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total < 1 {
		t.Fatal(errors.New("至少应存在一个满足条件的LigalIssue"))
	}
}

func testLigalIssueCategoryList(t *testing.T) {
	list, err := LegalIssueDao.CategoryList()
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
		UserId:  TestUserId,
	}
	err := LegalIssueFavoriteDao.Insert(issueFavorite)
	if err != nil {
		t.Fatal(err)
	}
	has, err := LegalIssueFavoriteDao.Exist(issueId, TestUserId)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Fatal(errors.New("LegalIssue应该已添加收藏"))
	}
	err = LegalIssueFavoriteDao.Delete(issueId, TestUserId)
	if err != nil {
		t.Fatal(err)
	}
	has, err = LegalIssueFavoriteDao.Exist(issueId, TestUserId)
	if err != nil {
		t.Fatal(err)
	}
	if has {
		t.Fatal(errors.New("LegalIssue应该已删除收藏"))
	}
}
