package model

import (
	"fmt"
	"law/conf"
	"log"
	"regexp"
	"strings"
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
var Db1 *xorm.Engine

// MySQL链接字符串
func Dsn(config *conf.Config) string {
	// 用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
	const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"
	return fmt.Sprintf(_dsn, config.Db.DbUser, config.Db.DbPasswd, config.Db.DbHost, config.Db.DbPort, config.Db.DbName, config.Db.DbParams)
}

func Init() {
	// 初始化数据库操作的 Xorm
	//连接到mysql
	//后面加入配置文件
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

// Page 分页基本数据
type Page struct {
	PageIndex int      `json:"page_index" form:"page_index" query:"page_index" validate:"gt=0"` //分页页码
	ItemNum   int      `json:"item_num" form:"item_num" query:"item_num" validate:"gt=0"`       //分页大小
	Order     []string `json:"order_arr[]" form:"order_arr[]" query:"order_arr[]"`
}

//处理分页及排序逻辑,page如:map[item_num:[8] order[]:[asc-assent_count asc-reply_count] page_index:[1]]
func (page *Page) GetResults(sess *xorm.Session, modsPtr interface{}, condiBean ...interface{}) (*PageResult, error) {
	defer sess.Close()
	sess.Limit(page.ItemNum, (page.PageIndex-1)*page.ItemNum) //分页
	orders := map[string]bool{"asc": true, "desc": true}
	regexpCompiled := regexp.MustCompile(`^[a-z,A-Z,0-9,\.,_]+$`) //正则表达式
	for _, item := range page.Order {
		arr := strings.Split(item, "-")
		if len(arr) != 2 || !orders[arr[0]] {
			return nil, fmt.Errorf("排序指令不合法:%s", item)
		}
		if !regexpCompiled.MatchString(arr[1]) {
			//由于直接使用了用户输入,所以要注意sql注入问题
			return nil, fmt.Errorf("检测到不合法字符:%s", arr[1])
		}
		if arr[0] == "asc" {
			sess.Asc(arr[1])
		} else {
			sess.Desc(arr[1])
		}
	}
	count, err := sess.FindAndCount(modsPtr)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		modsPtr = &[0]int{}
	}
	return &PageResult{Rows: modsPtr, Total: count}, nil
}

type PageResult struct {
	Rows  interface{} `json:"rows"`
	Total int64       `json:"total"`
}
