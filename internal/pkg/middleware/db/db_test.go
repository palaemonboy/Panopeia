package db

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

var gormConn = &gorm.DB{Config: &gorm.Config{}}

func TestUnit_Initializes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbConfigs := []DBConfig{
		{
			Name:        "xxx",
			User:        "xxx",
			Password:    "xxx",
			IP:          "127.0.0.1",
			Port:        "3306",
			DBName:      "xxx",
			Charset:     "utf8",
			MaxIdle:     "2",
			MaxOpen:     "10",
			MaxLiftTime: "1000",
		},
	}
	convey.Convey("TestUnit_Initializes", t, func() {
		convey.Convey("Err-01", func() {
			patches1 := gomonkey.ApplyFunc(gorm.Open, func(dialector gorm.Dialector, opts ...gorm.Option) (
				db *gorm.DB, err error) {
				return nil, errors.New("err")
			})
			defer patches1.Reset()
			err := Initializes(dbConfigs)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Err-02", func() {
			patches1 := gomonkey.ApplyFunc(gorm.Open, func(dialector gorm.Dialector, opts ...gorm.Option) (
				db *gorm.DB, err error) {
				return gormConn, nil
			})
			defer patches1.Reset()
			patches2 := gomonkey.ApplyFunc(gormConn.DB, func() (*sql.DB, error) {
				return nil, errors.New("err")
			})
			defer patches2.Reset()
			err := Initializes(dbConfigs)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Suc", func() {
			patches1 := gomonkey.ApplyFunc(gorm.Open, func(dialector gorm.Dialector, opts ...gorm.Option) (
				db *gorm.DB, err error) {
				return gormConn, nil
			})
			defer patches1.Reset()
			patches2 := gomonkey.ApplyMethod(reflect.TypeOf(gormConn), "DB", func(db *gorm.DB) (*sql.DB, error) {
				return &sql.DB{}, nil
			})
			defer patches2.Reset()
			err := Initializes(dbConfigs)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestUnit_GetTestDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_GetTestDB", t, func() {
		convey.Convey("Suc", func() {
			dbManager.ConnPool = make(map[string]*gorm.DB)
			dbManager.ConnPool[TestDB] = &gorm.DB{}
			got, err := GetTestDB()
			convey.So(err, convey.ShouldBeNil)
			convey.So(reflect.DeepEqual(got, &gorm.DB{}), convey.ShouldBeTrue)
		})
		convey.Convey("Err-01", func() {
			dbManager.ConnPool = make(map[string]*gorm.DB)
			got, err := GetTestDB()
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(reflect.DeepEqual(got, &gorm.DB{}), convey.ShouldBeFalse)
		})
		convey.Convey("Err-02", func() {
			dbManager.ConnPool = nil
			_, err := GetTestDB()
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(err.Error(), convey.ShouldEqual, "db has not initializes")
		})
	})
}
