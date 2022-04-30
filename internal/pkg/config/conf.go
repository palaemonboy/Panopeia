package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/palaemonboy/Panopeia/internal/pkg/middleware/db"
	"github.com/pkg/errors"
)

type Config struct {
	DB []db.DBConfig `yaml:"DB"`
}

// Initialize 初始化配置
func Initialize() (Config, error) {
	var config Config
	// 获取yaml文件
	configFile, err := ioutil.ReadFile("internal/pkg/config/config.yaml")
	if err != nil {
		return config, errors.Wrapf(err, "read config.yaml failed")
	}
	// unmarshal具体配置
	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return config, errors.Wrapf(err, "unmarshal config failed")
	}
	return config, nil
}
