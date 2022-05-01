package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

// 错误码划分，参照http状态码设计
// 分两段，xxx xxx
// 第一段为http状态码
// 第二段为具体定义原因
// example:
// 400001 表示参数错误，具体错误码为001

const (
	ParamError = 400001

	UnknownError = 500000
)

// Error 自定义携带错误码的error类型
type Error struct {
	Message string
	Code    int
}

// Error 返回下层函数传递的错误信息
func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, message: %v", e.Code, e.Message)
}

// New 生成error实例
func New(code int, message string) error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

// NewBadRequest 请求参数错误
func NewBadRequest(message string) error {
	return &Error{
		Message: message,
		Code:    ParamError,
	}
}

// RetErr 解析错误
func RetErr(err error) (errCode int, message string) {
	err = errors.Cause(err)
	switch err.(type) {
	case nil:
		return 0, ""
	case *Error:
		return err.(*Error).Code, err.(*Error).Message
	default:
		return UnknownError, err.Error()
	}
}
