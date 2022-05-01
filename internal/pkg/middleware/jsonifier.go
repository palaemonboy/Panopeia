package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/internal/pkg/errs"
)

const (
	contextKeyResp = "resp"
	contextKeyErr  = "errs"
)

// Response 所有Resp按照固定格式返回
type Response struct {
	Status  int         `json:"status"`
	Version string      `json:"version"`
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	TraceBack  string `json:"traceback"`
}

func Jsonifier(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// process request
		c.Next()

		shouldJsonify := false

		resp := &Response{
			Version: version,
		}

		statusCode := http.StatusOK

		if value, exists := c.Get(contextKeyResp); exists {
			resp.Status = 0
			resp.Success = true
			resp.Data = value
			resp.Error = nil
			shouldJsonify = true
		}

		if value, exists := c.Get(contextKeyErr); exists {
			if err, ok := value.(*errs.Error); ok {
				if err.Code != 0 {
					statusCode = err.Code
				}
			}

			resp.Success = false
			resp.Data = nil
			resp.Error = value
			shouldJsonify = true
		}

		if shouldJsonify {
			c.JSON(statusCode, resp)
		}
	}
}

func SetResp(c *gin.Context, value interface{}) {
	c.Set(contextKeyResp, value)
}

func SetErr(c *gin.Context, err error) {
	statusCode, message := errs.RetErr(err)
	c.Set(contextKeyErr, &Error{
		StatusCode: statusCode,
		Message:    message,
		TraceBack:  "",
	})
}

func SetErrWithTraceBack(c *gin.Context, err error) {
	statusCode, message := errs.RetErr(err)
	c.Set(contextKeyErr, &Error{
		StatusCode: statusCode,
		Message:    message,
		TraceBack:  fmt.Sprintf("%+v", err),
	})
}
