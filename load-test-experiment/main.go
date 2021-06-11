package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/light", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Light operations",
		})
	})

	r.GET("/medium", func(ctx *gin.Context) {
		time.Sleep(2 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Medium operations",
		})
	})

	r.GET("/heavy", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Heavy operations",
		})
	})

	s := &http.Server{
		Addr:           ":9000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}