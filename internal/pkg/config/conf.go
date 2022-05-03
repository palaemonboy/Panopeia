package config

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
	Log       *LogConfig
	DB        []DBConfig
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// DBConfig DB配置结构体
type DBConfig struct {
	Name        string `mapstructure:"name"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	IP          string `mapstructure:"ip"`
	Port        string `mapstructure:"port"`
	DBName      string `mapstructure:"dbName"`
	Charset     string `mapstructure:"charset"`
	MaxIdle     int    `mapstructure:"maxIdle"`
	MaxOpen     int    `mapstructure:"maxOpen"`
	MaxLiftTime int    `mapstructure:"maxLiftTime"`
}

func Init() (AppConfig, error) {
	var config AppConfig
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return config, errors.Wrapf(err, "init settings failed")
	}
	// 将读取的配置信息保存至config中
	if err := viper.Unmarshal(&config); err != nil {
		return config, errors.Wrapf(err, "Unmarshal settings failed")
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件更改成功")
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("Unmarshal settings failed, err:%v\n", err)
			return
		}
	})
	return config, nil
}
