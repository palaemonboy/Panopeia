package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/internal/pkg/errs"
	"github.com/palaemonboy/Panopeia/internal/pkg/middleware"
)

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {

	var req GetUsersReq
	var reqjson GetUsersReqJson
	//绑定 URL参数
	if err := c.ShouldBindQuery(&req); err != nil {
		middleware.SetErrWithTraceBack(c,
			errs.New(errs.ParamError, "invalid req"),
		)
		return
	}
	//绑定 Json参数
	if err := c.ShouldBindJSON(&reqjson); err != nil {
		middleware.SetErrWithTraceBack(c,
			errs.New(errs.ParamError, "invalid req"),
		)
		return
		//panic(err)
	}
	var resp GetUserResp
	resp.UserName = req.UserName
	resp.Message = reqjson.Message

	middleware.SetResp(c, resp)
}
