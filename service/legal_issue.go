package service

import (
	"errors"
	"law/model"
)

type legalIssueSrv struct{}

var LegalIssueSrv *legalIssueSrv = &legalIssueSrv{}

type legalIssueInfo struct {
	Id             int    `json:"id"`
	CreatorUid     int    `json:"creator_uid"`
	FirstCategory  string `json:"first_category"`
	SecondCategory string `json:"second_category"`
	Tags           string `json:"tags"`
	Title          string `json:"title"`
	Imgs           string `json:"imgs"`
	Content        string `json:"content"`
	CreateTime     int    `json:"create_time"`
}

func (l *legalIssueSrv) GetLegalIssue(legalIssueId int) (issueInfo *legalIssueInfo, err error) {
	issue := &model.LegalIssue{Id: legalIssueId}
	issue.Get()
	has, err := issue.Get()
	if err != nil {
		return
	}
	if !has {
		err = errors.New("此legalIssue不存在。")
		return
	}
	issueInfo = &legalIssueInfo{
		Id:             issue.Id,
		CreatorUid:     issue.CreatorUid,
		FirstCategory:  issue.FirstCategory,
		SecondCategory: issue.SecondCategory,
		Tags:           issue.Tags,
		Title:          issue.Title,
		Imgs:           issue.Imgs,
		Content:        issue.Content,
		CreateTime:     issue.CreateTime,
	}
	return
}

func (l *legalIssueSrv) AddFavorite(uid int, issueId int) (err error) {
	favorite := &model.LegalIssueFavorite{
		IssueId: issueId,
		UserId:  uid,
	}
	has, err := favorite.Exist()
	if err != nil {
		return
	}
	if has {
		return nil
	}
	err = favorite.Insert()
	return
}

func (l *legalIssueSrv) CancelFavorite(uid int, issueId int) (err error) {
	favorite := &model.LegalIssueFavorite{
		IssueId: issueId,
		UserId:  uid,
	}
	err = favorite.Delete()
	return
}

func (l *legalIssueSrv) IsFavorite(uid int, issueId int) (has bool, err error) {
	favorite := &model.LegalIssueFavorite{
		IssueId: issueId,
		UserId:  uid,
	}
	has, err = favorite.Exist()
	return
}
