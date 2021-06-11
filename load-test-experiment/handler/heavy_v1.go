package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type heavyV1Handler struct {}

type HeavyV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewHeavyV1Handler() HeavyV1Handler {
	return &heavyV1Handler{}
}

func (h *heavyV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Heavy operations",
		})
	}
}