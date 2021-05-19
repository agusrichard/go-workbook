package middlewares

import (
	"net/http"
	"twit/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header["Authorization"][0]

		user, err := utils.ParseToken(authorizationHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "You are unauthorized",
				"data":    nil,
			})
		} else {
			ctx.Set("UserID", user.ID)
			ctx.Set("Email", user.Email)
			ctx.Next()
		}
	}
}
