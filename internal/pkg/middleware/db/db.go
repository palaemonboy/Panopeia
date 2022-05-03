package db

import (
	"github.com/palaemonboy/Panopeia/internal/pkg/config"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// TestDB 测试DB
	TestDB = "test"
)

type GormManager struct {
	ConnPool map[string]*gorm.DB
}

var dbManager GormManager

// Initializes DB初始化
func Initializes(dbConfigs *config.DBConfig) error {
	dbManager.ConnPool = make(map[string]*gorm.DB)
	for _, dbConfig := range dbConfigs {
		conn, err := openGorm(dbConfig)
		if err != nil {
			return errors.Wrapf(err, "init %s db failed", dbConfig.Name)
		}
		dbManager.ConnPool[dbConfig.Name] = conn
	}
	return nil
}

// openGorm 根据DB配置初始化连接
func openGorm(config config.DBConfig) (*gorm.DB, error) {
	dsn := config.User + ":" + config.Password + "@tcp(" + config.IP + ":" + config.Port + ")/" + config.DBName +
		"?parseTime=True&loc=Local&charset=" + config.Charset
	if config.Charset == "" {
		dsn += "utf8"
	}
	gormConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrapf(err, "gorm open %v failed", config.Name)
	}
	sqlDB, err := gormConn.DB()
	if err != nil {
		return nil, errors.Wrapf(err, "gorm get sqlDB failed")
	}
	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MaxOpen)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxLiftTime) * time.Millisecond)
	return gormConn, nil
}

// GetTestDB 获取测试DB连接
func GetTestDB() (*gorm.DB, error) {
	if dbManager.ConnPool == nil {
		return nil, errors.Errorf("db has not initializes")
	}
	conn, ok := dbManager.ConnPool[TestDB]
	if !ok {
		return nil, errors.Errorf("%v has not initializes", TestDB)
	}
	return conn, nil
}
