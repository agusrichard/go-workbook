package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi-tested-app/entities"
)

type appHandler func(ctx *gin.Context) *entities.AppResult

func ServeHTTP(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx)
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, entities.Response{
				Success: false,
				Message: "INTERNAL SERVER ERROR",
				Data: nil,
			})
		}

		if result.Err == nil {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: true,
				Message: result.Message,
				Data: result.Data,
			})
		} else {
			ctx.JSON(result.StatusCode, entities.Response{
				Success: false,
				Message: result.Err.Error(),
				Data: result.Data,
			})
		}
	}
}