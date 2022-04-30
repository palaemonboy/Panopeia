package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	// 无参数
	r.GET("/start", func(context *gin.Context) {
		context.String(200, "Hello Word!")
	})

	_ = r.Run(":9090")
}
