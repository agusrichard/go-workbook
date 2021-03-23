package middlewares

import (
	"net/http"
	"strings"
	"twit/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header["Authorization"][0]
		token := strings.Split(authorizationHeader, " ")[1]

		user, err := utils.ParseToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
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
