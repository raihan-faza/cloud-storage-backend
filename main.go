package main

import (
	"cloud/app/service/auth/api"
	"cloud/cmd"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	cmd.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/hi", func(ctx *gin.Context) {
		hi := cmd.Test()
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": hi,
			},
		)
	})
	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.LoginUser)

	r.Run()
}
