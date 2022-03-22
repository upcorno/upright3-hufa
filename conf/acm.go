package conf

import (
	"math/rand"
	"os"
	"syscall"
	"time"

	zlog "github.com/rs/zerolog/log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var acmConfigFile string = "./acm.toml"

func getAndListenConfig(c chan<- os.Signal) {
	acmConfig := getAcmConfig()
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{ClientConfig: &constant.ClientConfig{
		Endpoint:       acmConfig.ClientConfig.Endpoint,
		NamespaceId:    acmConfig.ClientConfig.NamespaceId,
		AccessKey:      acmConfig.ClientConfig.AccessKey,
		SecretKey:      acmConfig.ClientConfig.SecretKey,
		CacheDir:       acmConfig.ClientConfig.CacheDir,
		LogDir:         acmConfig.ClientConfig.LogDir,
		LogLevel:       acmConfig.ClientConfig.LogLevel,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}})
	if err != nil {
		zlog.Fatal().Msgf("configClient init fail.err:%s", err.Error())
		return
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: acmConfig.ConfigParam.DataId,
		Group:  acmConfig.ConfigParam.Group,
	})
	zlog.Info().Msgf("%s 获取到最新配置文件", defConfigFile)
	if err != nil {
		zlog.Fatal().Msgf("get config content fail.err:%s", err.Error())
		return
	}
	putContentToFile(&content)
	// 监听配置
	configClient.ListenConfig(vo.ConfigParam{
		DataId: acmConfig.ConfigParam.DataId,
		Group:  acmConfig.ConfigParam.Group,
		OnChange: func(namespace, group, dataId, data string) {
			zlog.Info().Msgf("%s 配置文件发生变化", defConfigFile)
			//随机时间后,发送信号,使进程退出,待重启时更新配置文件
			rand.Seed(time.Now().UnixNano())
			delayTime := rand.Intn(30) * (int)(time.Second)
			time.Sleep(time.Duration(delayTime))
			c <- syscall.SIGINT
		},
	})
}

func getAcmConfig() *acmConfig {
	viper.SetConfigFile(acmConfigFile)
	if err := viper.ReadInConfig(); err != nil {
		zlog.Fatal().Msgf("acm config init while viper.ReadInConfig error.err:%s", err.Error())
	}
	cfg := &acmConfig{}
	if err := viper.Unmarshal(cfg); err != nil {
		zlog.Fatal().Msgf("acm config init while viper.Unmarshal error.err:%s", err.Error())
	}
	return cfg
}

func putContentToFile(content *string) {
	os.WriteFile(defConfigFile, []byte(*content), os.FileMode(0666))
}

type acmConfig struct {
	ClientConfig *clientConfig `mapstructure:"client"`
	ConfigParam  *configParam  `mapstructure:"param"`
}

type clientConfig struct {
	Endpoint    string `mapstructure:"endpoint"`
	NamespaceId string `mapstructure:"namespace_id"`
	AccessKey   string `mapstructure:"access_key"`
	SecretKey   string `mapstructure:"secret_key"`
	CacheDir    string `mapstructure:"cache_dir"`
	LogDir      string `mapstructure:"log_dir"`
	LogLevel    string `mapstructure:"log_level"`
}
type configParam struct {
	DataId string `mapstructure:"data_id"`
	Group  string `mapstructure:"group"`
}
