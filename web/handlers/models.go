package handlers

type GetUsersReq struct {
	UserName string `form:"username" binding:"required"`
}
