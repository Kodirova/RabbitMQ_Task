package api

import (
	"test_task/storage/api/handler"
	"test_task/storage/internal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	Storage internal.Storage
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlers := handler.New(&handler.HandlerOptions{
		Storage: opt.Storage,
	})
	apis := router.Group("/storage")

	{
		apis.GET("/phone/:id", handlers.GetPhone)
	}
	return router
}
