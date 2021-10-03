package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, world!")
}
