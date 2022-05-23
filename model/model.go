package model

import (
	"fmt"
	"law/conf"
	_ "law/utils"
	"log"
	"time"

	zlog "github.com/rs/zerolog/log"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"

	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
	xlog "xorm.io/xorm/log"
)

// Db 数据库操作句柄
var Db *xorm.Engine

// MySQL链接字符串
func Dsn(config *conf.Config) string {
	// 用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
	const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"
	return fmt.Sprintf(_dsn, config.Db.DbUser, config.Db.DbPasswd, config.Db.DbHost, config.Db.DbPort, config.Db.DbName, config.Db.DbParams)
}

func init() {
	db, err := xorm.NewEngine("mysql", Dsn(conf.App))
	if err != nil {
		zlog.Fatal().Msgf("xorm.NewEngine初始化失败.err:%s", err.Error())
	}
	if conf.App.Orm.OrmHijackLog {
		sl := &xlog.SimpleLogger{
			DEBUG: log.New(zlog.Logger, "", 0),
			ERR:   log.New(zlog.Logger, "", 0),
			INFO:  log.New(zlog.Logger, "", 0),
			WARN:  log.New(zlog.Logger, "", 0),
		}
		if conf.App.IsDev() {
			sl.SetLevel(xlog.LOG_DEBUG)
		} else {
			sl.SetLevel(xlog.LOG_WARNING)
		}
		db.SetLogger(sl)
	}
	if err = db.Ping(); err != nil {
		zlog.Fatal().Msgf("数据库 ping失败.err:%s", err.Error())
	}
	db.DatabaseTZ = time.Local
	db.TZLocation = time.Local
	db.SetMaxIdleConns(conf.App.Orm.OrmIdle)
	db.SetMaxOpenConns(conf.App.Orm.OrmOpen)
	db.SetConnMaxLifetime(time.Duration(conf.App.Orm.OrmConnMaxLifetime) * time.Second)
	db.ShowSQL(conf.App.Orm.OrmShow)
	if conf.App.Orm.OrmCacheUse {
		cacher := caches.NewLRUCacher(caches.NewMemoryStore(), conf.App.Orm.OrmCacheSize)
		db.SetDefaultCacher(cacher)
	}
	if conf.App.Orm.OrmSync {
		err := db.Sync2(
			new(User),
			new(LegalIssue),
			new(LegalIssueFavorite),
			new(Consultation),
			new(ConsultationReply),
			new(InfringementMonitor),
			new(RightsProtection),
		)
		if err != nil {
			zlog.Fatal().Msgf("数据库 sync失败.err:%s", err.Error())
		}
	}
	_, err = db.Exec("ALTER TABLE legal_issue ADD FULLTEXT INDEX search_text (search_text) WITH PARSER ngram")
	if err != nil {
		zlog.Fatal().Msgf("创建全文索引失败：.err:%s", err.Error())
	}
	Db = db
	zlog.Info().Msg("model init")
}

func CountNewItems(minId int, table string) (count int, maxId int, err error) {
	ids := &[]map[string]int{}
	err = Db.Table(table).
		Cols("id").
		Where("id > ?", minId).
		Desc("id").
		Find(ids)
	if err != nil {
		return
	}
	count = len(*ids)
	if count > 0 {
		maxId = (*ids)[0]["id"]
	} else {
		maxId = minId
	}
	return
}
