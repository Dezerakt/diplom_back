package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseUtil struct {
}

func ErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	ctx.Abort()
}
