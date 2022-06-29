package service

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"law/conf"
	"net/http"
	"sync"
	"time"

	"github.com/eko/gocache/v3/marshaler"
	"github.com/eko/gocache/v3/store"
)

//放置不便于归类的服务

type mixed struct{}

var MixedSvr *mixed

type fundUsage struct {
	UpdateTimeStr   string
	ObligeeProvince string
	Amount          int
	PaymentNum      int
	CaseNum         int
}

var usageCache marshaler.Marshaler
var usageMutex sync.Mutex

func init() {
	usageCache = *marshaler.New(store.NewRedis(Rdb))
}

//查询维权基金的使用情况
func (c *mixed) GetFundUsage(obligeeProvince string) (usage *fundUsage, err error) {
	usage, err = c.requestFundUsage(obligeeProvince)
	h := md5.New()
	h.Write([]byte(conf.App.ProjectName + "GetFundUsage" + obligeeProvince))
	key := fmt.Sprintf("%x", h.Sum(nil))
	_, err = usageCache.Get(context.Background(), key, &usage)
	if err == nil {
		return
	}
	usageMutex.Lock()
	defer usageMutex.Unlock()
	//锁控制后再次检查缓存，减少从源端查询次数
	_, err = usageCache.Get(context.Background(), key, &usage)
	if err == nil {
		return
	}
	usage, err = c.requestFundUsage(obligeeProvince)
	if err == nil {
		t := time.Now()
		dur := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local).Sub(t)
		err = usageCache.Set(context.Background(), key, usage, store.WithExpiration(dur))
		return
	}
	return
}

func (c *mixed) requestFundUsage(obligeeProvince string) (usage *fundUsage, err error) {
	url := conf.App.Mixed.FundUsageUrl + "?obligee_province=" + obligeeProvince
	client := http.DefaultClient
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body := string(bytes)
	tmp := &struct {
		Code int `json:"code"`
		Data struct {
			CaseNum                    int `json:"case_num"`
			CaseExpenditureNum         int `json:"case_expenditure_num"`
			CaseExpenditureTotalAmount int `json:"case_expenditure_total_amount"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(bytes, tmp)
	if err != nil || tmp.Code != 0 {
		err = fmt.Errorf("查询维权基金的使用情况失败。 body: %#v ,err:%#v ", body, err)
		return
	}
	usage = &fundUsage{
		ObligeeProvince: obligeeProvince,
		UpdateTimeStr:   time.Now().Format("2006年01月02日"),
		Amount:          tmp.Data.CaseExpenditureTotalAmount,
		CaseNum:         tmp.Data.CaseNum,
		PaymentNum:      tmp.Data.CaseExpenditureNum,
	}
	return
}
