package handlers

type GetUsersReq struct {
	UserName string `form:"username" binding:"required"`
}
type GetUsersReqJson struct {
	Message string `json:"message" binding:"required"`
}

type GetUserResp struct {
	UserName string `json:"username"`
	Message  string `json:"message"`
}
