package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/internal/pkg/errs"
	"github.com/palaemonboy/Panopeia/internal/pkg/middleware"
)

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {

	var req GetUsersReq

	if err := c.ShouldBindQuery(&req); err != nil {
		middleware.SetErrWithTraceBack(c,
			errs.New(errs.ParamError, "invalid req"),
		)
		return
	}
	Mesaage := "Get All users haha."
	var resp GetUserResp
	resp.UserName = req.UserName
	resp.Message = Mesaage

	middleware.SetResp(c, resp)
}
