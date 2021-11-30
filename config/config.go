package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 初始化配置文件
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 有指定配置文件，则解析配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 没有指定则解析默认配置文件 config/config.yaml
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	// 设置配置文件格式为yaml
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 设置读取环境变量的前缀,Viper可以从环境变量读取配置
	// 如环境变量中的 export APISERVER_ADDR=:7777
	// export APISERVER_URL=http://127.0.0.1:7777
	// 重新启动服务后，此时监听7777 此时环境变量会映射成yaml中的变量，并覆盖掉配置中的值
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_") // 环境变量用_连接，会映射成yaml中的变量名
	viper.SetEnvKeyReplacer(replacer)

	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file change: %s", e.Name)
	})
}
