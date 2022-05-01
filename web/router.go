package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/palaemonboy/Panopeia/internal/pkg/middleware/db"

	"github.com/palaemonboy/Panopeia/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/internal/pkg/middleware"
	"github.com/palaemonboy/Panopeia/web/handlers"
)

// Router 路由结构体
type Router struct {
	*gin.Engine
}

// NewRouter 代码入口
func NewRouter() *Router {
	//  配置初始化
	conf, err := config.Initialize()
	if err != nil {
		// 前置配置加载失败，直接panic
		panic(err)
	}
	// DB初始化
	if err := db.Initializes(conf.DB); err != nil {
		// 前置DB初始化失败，直接panic
		panic(err)
	}
	// 获取DB连接，测试数据库使用
	_, err = db.GetTestDB()
	router := gin.Default()
	version := "0.0.1"
	router.Use(
		middleware.Jsonifier(version),
	)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Panopeia Server.")
	})
	router.GET("/start", func(c *gin.Context) {
		middleware.SetResp(c, "Panopeia Running.")
	})
	// API入口
	API := router.Group("/apis")
	// Service 入口
	ServiceAPI := API.Group("/service")
	{
		// Login 入口
		LoginAPI := ServiceAPI.Group("/login")
		{
			LoginAPI.GET("/getusers", handlers.GetUsers)
		}

	}

	return &Router{
		Engine: router,
	}
}

// Run 监听连接
func (r *Router) Run() {

	srv := &http.Server{
		Addr:    ":9090",
		Handler: r.Engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}

// Run api入口
func Run() {
	NewRouter().Run()
}
