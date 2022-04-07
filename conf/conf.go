package conf

import (
	"flag"
	"os"
	"time"

	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	App           *Config           //运行配置实体
	defConfigFile = "./config.toml" //配置文件路径，方便测试
)

type Config struct {
	Mode      string         `mapstructure:"mode"`
	Jwt       *jwtConf       `mapstructure:"jwt"`
	Http      *httpConf      `mapstructure:"http"`
	Orm       *ormConf       `mapstructure:"orm"`
	Db        *dbConf        `mapstructure:"db"`
	Rpc       *rpcConf       `mapstructure:"rpc"`
	Ristretto *ristrettoConf `mapstructure:"ristretto"`
	Oss       *oss           `mapstructure:"oss"`
	WxApp     *wxApp         `mapstructure:"wx_app"`
	Account   *accountConf   `mapstructure:"account"`
}

func NewConfig() *Config {
	return &Config{
		Http: &httpConf{
			Address:           "",
			ReadTimeout:       20 * time.Second,
			WriteTimeout:      20 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			IdleTimeout:       10 * time.Second,
		},
		Ristretto: &ristrettoConf{
			NumCounters: 5000000,
			MaxCost:     20000000,
			BufferItems: 64,
		},
	}
}

func Init(c chan<- os.Signal) {
	var cfgFile string
	// 从启动命令中读取配置文件路径
	flag.StringVar(&cfgFile, "c", defConfigFile, "path of config file.")
	flag.Parse()
	if cfgFile == defConfigFile && fileExist(acmConfigFile) {
		getAndListenConfig(c)
	}
	if cfgFile == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	} else {
		viper.SetConfigFile(cfgFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		zlog.Fatal().Msgf("config init while viper.ReadInConfig error.err:%s", err.Error())
	}
	cfg := NewConfig()
	if err := viper.Unmarshal(cfg); err != nil {
		zlog.Fatal().Msgf("config init while viper.Unmarshal error.err:%s", err.Error())
	}
	App = cfg
}
func (app *Config) IsProd() bool {
	return app.Mode == "prod"
}
func (app *Config) IsDev() bool {
	return app.Mode == "dev"
}
func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// jwt config
type jwtConf struct {
	LoginKey       string `mapstructure:"login_key"`
	LoginPath      string `mapstructure:"login_path"`
	BackgroundPath string `mapstructure:"background_path"`
	AuthKey        string `mapstructure:"auth_key"`
	AuthLifetime   int    `mapstructure:"auth_lifetime"`
}

// http config
type httpConf struct {
	Address           string        `mapstructure:"address"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
}

// orm config
type ormConf struct {
	OrmIdle      int  `mapstructure:"orm_idle"`       //
	OrmOpen      int  `mapstructure:"orm_open"`       //
	OrmShow      bool `mapstructure:"orm_show"`       //显示sql
	OrmSync      bool `mapstructure:"orm_sync"`       //同步表结构
	OrmCacheUse  bool `mapstructure:"orm_cache_use"`  //是否使用缓存
	OrmCacheSize int  `mapstructure:"orm_cache_size"` //缓存数量
	OrmHijackLog bool `mapstructure:"orm_hijack_log"` //劫持日志
}

// db config
type dbConf struct {
	DbHost    string `mapstructure:"db_host"`     //数据库地址
	DbPort    int    `mapstructure:"db_port"`     //数据库端口
	DbUser    string `mapstructure:"db_user"`     //数据库账号
	DbPasswd  string `mapstructure:"db_passwd"`   //数据库密码
	DbName    string `mapstructure:"db_name"`     //数据库名称
	DbParams  string `mapstructure:"db_params"`   //数据库参数
	YsjDbName string `mapstructure:"ysj_db_name"` //数据库名称
}

// rpcConf config
type rpcConf struct {
	MessageQueryPushUrl string `mapstructure:"message_query_push_url"` //消息队列服务地址
	Secret              string `mapstructure:"secret"`                 //请求秘钥
}

//ristretto cache config
type ristrettoConf struct {
	NumCounters int64 `mapstructure:"num_counters"`
	MaxCost     int64 `mapstructure:"max_cost"`
	BufferItems int64 `mapstructure:"buffer_items"`
}

type oss struct {
	RegionId        string `mapstructure:"region_id"`
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	RoleArn         string `mapstructure:"role_arn"`
	BucketName      string `mapstructure:"bucket_name"`
	Host            string `mapstructure:"host"`
}

type wxApp struct {
	Appid  string `mapstructure:"appid"`
	Secret string `mapstructure:"secret"`
}

type accountConf struct {
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
}
