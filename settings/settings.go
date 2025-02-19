package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	// 1. 接受参数
	viper.SetConfigFile(filePath)
	
	// 2. 直接获取路径
	// viper.SetConfigFile("./config.yaml")

	// 3. 指定配置文件名和路径，指定配置文件名（不带后缀），指定查找配置文件的路径
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config")

	// 4. 配合远程配置中心使用，告诉 viper 当前的数据使用什么格式去解析
	// viper.SetConfigType("json")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config files have been changed...")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
			return
		}
	})
	
	return
}
