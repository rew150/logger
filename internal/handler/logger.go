package handler

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(
			http.StatusBadRequest,
			"bad request: %s",
			err,
		)
		return
	}

	log := logger.Log{
		Timestamp: time.Now(),
		Message:   string(body),
	}

	if err := logger.Service.Append(ctx.Request.Context(), log); err != nil {
		ctx.String(http.StatusInternalServerError, "db error: %s", err)
		return
	}

	ctx.Status(http.StatusCreated)

}
