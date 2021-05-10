package api

import "github.com/gin-gonic/gin"

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(ctx *gin.Context) {
	resp := HelloResponse{
		Message: "Hello World",
	}
	ctx.JSON(200, resp)
}
