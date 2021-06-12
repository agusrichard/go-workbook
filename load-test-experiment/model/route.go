package model

import "github.com/gin-gonic/gin"

type Route struct {
	Action string
	Handler func(ctx *gin.Context)
}
