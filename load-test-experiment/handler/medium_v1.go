package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type mediumV1Handler struct {}

type MediumV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewMediumV1Handler() MediumV1Handler {
	return &mediumV1Handler{}
}

func (h *mediumV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(2 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Medium operations",
		})
	}
}