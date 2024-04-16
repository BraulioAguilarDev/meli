package http

import (
	"meli/internal/adapter/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	itemHandler ItemHdlr,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = strings.Split(config.AllowedOrigins, ",")

	router := gin.New()
	router.Use(gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("/upload-file", itemHandler.UploadFile)
		}
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(addr string) error {
	return r.Run(addr)
}
