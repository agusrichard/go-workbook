package utils

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func LogAbort(ctx *gin.Context, err error, httpStatus int) {
	if err != nil {
		log.Println(fmt.Sprintf("Error Message: %s\n", err.Error()))
		ctx.AbortWithStatusJSON(httpStatus, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	}
}
