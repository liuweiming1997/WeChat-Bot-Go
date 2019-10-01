package conf

import (
	"github.com/WeChat-Bot-Go/logger"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	REDIS_HOST string `default:"127.0.0.1:6379"`
}

var GlobalConfig *Config = &Config{}

func init() {
	// TODO(weimingliu) why can not print here
	// fmt.Sprintf("%T", GlobalConfig)
	err := envconfig.Process("ENV_PREFIX", GlobalConfig)
	if err != nil {
		logger.Fatal("Init env config error ", err)
	}
	logger.Info(GlobalConfig)
}
