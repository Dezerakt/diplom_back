package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUtil struct {
}

func ErrorResponse(ctx *gin.Context, err error) {

	ctx.JSON(http.StatusBadRequest, err.Error())
	ctx.Abort()
	return
}
