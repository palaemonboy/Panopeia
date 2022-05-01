package errs

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestUnit_New(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_New", t, func() {
		convey.Convey("Suc", func() {
			err := New(400001, "bad req")
			convey.So(err.Error(), convey.ShouldEqual, "Error: 400001, message: bad req")
		})
	})
}

func TestUnit_NewBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_NewBadRequest", t, func() {
		convey.Convey("Suc", func() {
			err := NewBadRequest("bad req")
			convey.So(err.Error(), convey.ShouldEqual, "Error: 400001, message: bad req")
		})
	})
}

func TestUnit_RetErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	convey.Convey("TestUnit_RetErr", t, func() {
		convey.Convey("Nil", func() {
			errCode, message := RetErr(nil)
			convey.So(errCode, convey.ShouldBeZeroValue)
			convey.So(message, convey.ShouldBeZeroValue)
		})
		convey.Convey("Error", func() {
			errCode, message := RetErr(NewBadRequest("err"))
			convey.So(errCode, convey.ShouldEqual, 400001)
			convey.So(message, convey.ShouldEqual, "err")
		})
		convey.Convey("Unknown", func() {
			errCode, message := RetErr(errors.New("err"))
			convey.So(errCode, convey.ShouldEqual, 500000)
			convey.So(message, convey.ShouldEqual, "err")
		})
	})
}
