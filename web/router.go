package web

import (
	"context"
	"fmt"
	"github.com/palaemonboy/Panopeia/internal/pkg/errs"
	"github.com/palaemonboy/Panopeia/internal/pkg/middleware/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// 1、加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		panic(err)
	}

	// 2、初始化日志
	if err := logger.Init(config.Conf.Log, config.Conf.Mode); err != nil {
		fmt.Printf("init log settings failed, err:%v\n", err)
		panic(err)
	}
	// 日志追加
	defer zap.L().Sync()

	//// 3、初始化DB连接
	//if err := db.Initializes(config.Conf.DB); err != nil {
	//	fmt.Printf("init settings failed, err:%v\n", err)
	//	panic(err)
	//}
	//defer db.Close()
	// 获取DB连接，测试数据库使用
	//_, err = db.GetTestDB()

	router := gin.Default()
	version := "0.0.1"
	router.Use(
		middleware.Jsonifier(version), // 格式化输出中间件
		middleware.CORS(),             //跨域中间件
		logger.GinLogger(),            //日志中间件
		logger.GinRecovery(true),      //Recovery 中间件
	)
	router.NoRoute(func(c *gin.Context) {
		middleware.SetErrWithTraceBack(c,
			errs.New(http.StatusNotFound, "Default 404."),
		)
	})
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
			LoginAPI.POST("/getusers", handlers.GetUsers)
		}

	}

	return &Router{
		Engine: router,
	}
}

// Run 监听连接
func (r *Router) Run() {

	srv := &http.Server{
		Addr:    ":9999",
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
