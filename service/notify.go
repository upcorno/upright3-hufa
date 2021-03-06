package service

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"law/conf"
	dao "law/dao"
	"law/enum"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	zlog "github.com/rs/zerolog/log"
)

type NotifySrv struct{}

var Notify *NotifySrv
var timeInterval int8 = 2
var (
	lastConsultationId, lastProtectId, lastMonitorId int
)

func init() {
	go func() {
		//不影响进程整体启动
		time.Sleep(time.Second * 2)
		//初始化启动时的lastConsultationId, lastProtectId, lastMonitorId
		_, _, _, err := countNewItems()
		if err != nil {
			panic(err)
		}
		for {
			time.Sleep(time.Duration(timeInterval) * time.Minute)
			go func() {
				Notify.NewBusinessNotifyByEmail()
			}()
		}
	}()
}

//NewToReplyBusiness 新业务邮件通知
func (n *NotifySrv) NewBusinessNotifyByEmail() {
	min := time.Now().Minute()
	min += min % int(timeInterval)
	key := string(md5.New().Sum([]byte(conf.App.ProjectName + "notify-BusinessToReplyByEmail" + fmt.Sprint(min))))
	ok, err := Rdb.SetNX(context.Background(), key, true, time.Duration(timeInterval)*time.Minute).Result()
	if err != nil {
		zlog.Error()
		return
	}
	if !ok {
		return
	}
	countConsultation, countProtect, countMonitor, err := countNewItems()
	if err != nil {
		zlog.Error()
		return
	}
	if countConsultation == 0 && countProtect == 0 && countMonitor == 0 {
		return
	}
	subject := fmt.Sprintf("【沪盾】有新的待回复内容。咨询： %d 个，维权： %d 个，监测： %d 个", countConsultation, countProtect, countMonitor)
	body := fmt.Sprintf("<html><body>您好！<br> %s <br> 后台地址： %s  </body></html>", subject, `<a class="resource_target" href="https://legal-consulting.youshangjiao.com.cn/" target="_blank">点击前往后台</a>`)
	sendEmail(subject, body, conf.App.Mail.NewBusinessReceivers...)
}

func countNewItems() (countConsultation int, countProtect int, countMonitor int, err error) {
	countConsultation, lastConsultationId, err = dao.ConsulDao.CountNewItems(lastConsultationId)
	if err != nil {
		return
	}
	countProtect, lastProtectId, err = dao.CooperationDao.CountNewItems(lastProtectId, enum.PROTECT)
	if err != nil {
		return
	}
	countMonitor, lastMonitorId, err = dao.CooperationDao.CountNewItems(lastMonitorId, enum.MONITOR)
	if err != nil {
		return
	}
	return
}

func sendEmail(subject string, body string, revivers ...string) {
	account := conf.App.Mail.Account
	password := conf.App.Mail.Password
	host := conf.App.Mail.Host
	auth := smtp.PlainAuth("", account, password, host)
	e := email.NewEmail()
	e.From = account
	e.To = revivers
	e.Subject = subject
	e.HTML = []byte(body)
	e.SendWithTLS(host+":465", auth, &tls.Config{ServerName: host})
}
