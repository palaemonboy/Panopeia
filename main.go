package main

import (
	"fmt"

	"github.com/palaemonboy/Panopeia/internal/pkg/middleware/db"

	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/internal/pkg/config"
)

func main() {
	// 初始化config配置
	conf, err := config.Initialize()
	if err != nil {
		panic(err)
	}
	fmt.Println(conf)
	// 初始化DB连接
	if err := db.Initializes(conf.DB); err != nil {
		panic(err)
	}
	// 获取测试DB的链接
	testDB, err := db.GetTestDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(testDB)
	r := gin.Default()
	// 无参数
	r.GET("/start", func(context *gin.Context) {
		context.String(200, "Hello Word!")
	})

	_ = r.Run(":9090")
}
