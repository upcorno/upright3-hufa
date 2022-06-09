package service

import (
	"fmt"
	dao "law/dao"
)

type legalIssueSrv struct{}

var LegalIssueSrv *legalIssueSrv = &legalIssueSrv{}

func (l *legalIssueSrv) List(page *dao.Page, search *dao.LegalIssueSearch) (issues *dao.PageResult, err error) {
	issues, err = dao.LegalIssueDao.List(page, search)
	return
}

func (l *legalIssueSrv) Get(legalIssueId int) (issue *dao.LegalIssue, err error) {
	issue, err = dao.LegalIssueDao.MustGet(legalIssueId)
	return
}

func (l *legalIssueSrv) Create(issue *dao.LegalIssue) (issueId int, err error) {
	issue.SearchText = fmt.Sprintf("%s%s%s%s%s%s", issue.BusinessCategory, issue.SecondCategory, issue.FirstCategory, issue.Tags, issue.Title, issue.Content)
	issueId, err = dao.LegalIssueDao.Insert(issue)
	return
}

func (l *legalIssueSrv) Update(issueId int, issue *dao.LegalIssue) (err error) {
	err = dao.LegalIssueDao.Update(issueId, issue, "first_category", "second_category", "business_category", "tags", "title", "imgs", "content")
	return
}

func (l *legalIssueSrv) Delete(issueId int) (err error) {
	err = dao.LegalIssueDao.Delete(issueId)
	return
}

func (l *legalIssueSrv) AddFavorite(uid int, issueId int) (err error) {
	has, err := dao.LegalIssueFavoriteDao.Exist(issueId, uid)
	if err != nil {
		return
	}
	if has {
		return nil
	}
	favorite := &dao.LegalIssueFavorite{
		IssueId: issueId,
		UserId:  uid,
	}
	err = dao.LegalIssueFavoriteDao.Insert(favorite)
	return
}

func (l *legalIssueSrv) CancelFavorite(uid int, issueId int) (err error) {
	err = dao.LegalIssueFavoriteDao.Delete(issueId, uid)
	return
}

func (l *legalIssueSrv) IsFavorite(uid int, issueId int) (has bool, err error) {
	has, err = dao.LegalIssueFavoriteDao.Exist(issueId, uid)
	return
}
