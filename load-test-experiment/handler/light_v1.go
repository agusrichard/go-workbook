package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type lightV1Handler struct {}

type LightV1Handler interface{
	Handle() gin.HandlerFunc
}

func NewLightV1Handler() LightV1Handler {
	return &lightV1Handler{}
}

func (h *lightV1Handler) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Light operations",
		})
	}
}