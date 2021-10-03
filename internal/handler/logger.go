package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rew150/logger/internal/logger"
)

func GetLogsHandler(ctx *gin.Context) {
	logs, err := logger.Service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusInternalServerError, "db error: %s", err)
		return
	}
	ctx.JSON(http.StatusOK, logs)
}

func AppendLogHandler(ctx *gin.Context) {
	var body logger.Log
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.String(
			http.StatusBadRequest,
			"error formatting body: %s",
			err,
		)
		return
	}

	if body.Timestamp == "" {
		ctx.String(
			http.StatusUnprocessableEntity,
			"timestamp is required",
		)
		return
	}

	if err := logger.Service.Append(ctx.Request.Context(), body); err != nil {
		ctx.String(http.StatusInternalServerError, "db error: %s", err)
		return
	}

	ctx.Status(http.StatusCreated)

}
