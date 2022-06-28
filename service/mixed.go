package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"law/conf"
	"net/http"
	"sync"
	"time"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/thedevsaddam/gojsonq/v2"
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

var usageCache *cache.Cache[fundUsage]
var usageMutex sync.Mutex

func init() {
	usageCache = cache.New[fundUsage](store.NewRedis(Rdb))
}

//查询维权基金的使用情况
func (c *mixed) GetFundUsage(obligeeProvince string) (usage *fundUsage, err error) {
	key := string(md5.New().Sum([]byte(conf.App.ProjectName + "GetFundUsage" + obligeeProvince)))
	cachedUsage, err := usageCache.Get(context.Background(), key)
	if err == nil {
		usage = &cachedUsage
		return
	}
	usageMutex.Lock()
	defer usageMutex.Unlock()
	//锁控制后再次检查缓存，减少从源端查询次数
	cachedUsage, err = usageCache.Get(context.Background(), key)
	if err == nil {
		usage = &cachedUsage
		return
	}
	usage, err = c.requestFundUsage(obligeeProvince)
	if err != nil {
		t := time.Now()
		dur := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local).Sub(t)
		usageCache.Set(context.Background(), key, *usage, store.WithExpiration(dur))
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
	jsonObj := gojsonq.New().FromString(body)
	codeV := jsonObj.Find("code")
	if code, ok := codeV.(float64); !ok || code != 0 {
		err = fmt.Errorf("查询维权基金的使用情况失败。 body: %#v ", body)
		return
	}
	usage = &fundUsage{ObligeeProvince: obligeeProvince,
		UpdateTimeStr: time.Now().Format("2006年01月02日")}
	jsonObj.Reset()
	caseNumV := jsonObj.Find("data.case_num")
	if caseNum, ok := caseNumV.(float64); !ok {
		err = fmt.Errorf("查询维权基金的使用情况-获取案件数失败。 body: %#v", body)
		return
	} else {
		usage.CaseNum = int(caseNum)
	}
	jsonObj.Reset()
	paymentNumV := jsonObj.Find("data.case_expenditure_num")
	if paymentNum, ok := paymentNumV.(float64); !ok {
		err = fmt.Errorf("查询维权基金的使用情况-获取支付笔数失败。 body: %#v", body)
		return
	} else {
		usage.PaymentNum = int(paymentNum)
	}
	jsonObj.Reset()
	amountV := jsonObj.Find("data.case_expenditure_total_amount")
	if amount, ok := amountV.(float64); !ok {
		err = fmt.Errorf("查询维权基金的使用情况-获取累计支付金额失败。 body: %#v", body)
		return
	} else {
		usage.Amount = int(amount)
	}
	return
}
