package config

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_Initialize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_Initialize", t, func() {
		convey.Convey("Err-01", func() {
			patches1 := gomonkey.ApplyFunc(ioutil.ReadFile, func(fileName string) ([]byte, error) {
				return nil, errors.New("err")
			})
			defer patches1.Reset()
			config, err := Initialize()
			convey.So(reflect.DeepEqual(config, Config{}), convey.ShouldEqual, true)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}
