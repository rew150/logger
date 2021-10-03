package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rew150/logger/internal/handler"
)

func Route(g *gin.Engine) {
	api := g.Group("/api")
	api.GET("/", handler.HelloHandler)
}
