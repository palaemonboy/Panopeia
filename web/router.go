package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/palaemonboy/Panopeia/pkg/middlewares"
	"github.com/palaemonboy/Panopeia/web/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Router struct {
	*gin.Engine
}

//  代码入口

func NewRouter() *Router {
	router := gin.Default()
	version := "0.0.1"
	router.Use(
		middlewares.Jsonifier(version),
	)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Panopeia Server.")
	})
	router.GET("/start", func(c *gin.Context) {
		middlewares.SetResp(c, "Panopeia Running.")
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
		router,
	}
}

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
func Run() {
	NewRouter().Run()
}
