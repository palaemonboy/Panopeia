package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/pkg/middlewares"
	"github.com/pkg/errors"
	"net/http"
)

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {

	var req GetUsersReq

	if err := c.ShouldBindQuery(&req); err != nil {
		middlewares.SetErrWithTraceBack(c,
			http.StatusNotFound,
			errors.Wrapf(err, "invalid req"),
		)
		return
	}
	Mesaage := "Get All users haha."
	var resp GetUserResp
	resp.UserName = req.UserName
	resp.Message = Mesaage

	middlewares.SetResp(c, resp)
}
