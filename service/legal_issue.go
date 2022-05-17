package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"law/model"
	"time"

	"github.com/eko/gocache/v2/store"
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

//由于问题列表接口内容几乎不会变化，因此进行了5min的缓存，
//如果请求参数指定根据是否收藏参数检索，则不会使用缓存
func (l *legalIssueSrv) LegalIssueList(page *model.Page, search *model.LegalIssueSearch) (issues *model.PageResult, err error) {
	if search.OnlyFavorite {
		issues, err = model.LegalIssueList(page, search)
		return
	}
	signStr := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", page, search))))
	ctx := context.TODO()
	cache, err := CacheManager.Get(ctx, signStr)
	if err == nil {
		issues = cache.(*model.PageResult)
		return
	}
	issues, err = model.LegalIssueList(page, search)
	if err == nil {
		CacheManager.Set(ctx, signStr, issues, &store.Options{Expiration: time.Second * 300})
	}
	return
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
