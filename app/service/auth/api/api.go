package api

import (
	"cloud/app/service/auth/controllers"
	"cloud/app/service/auth/middleware"
	"cloud/app/service/auth/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var request requests.RegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		panic(err)
	}
	token, err := controllers.CreateUser(request.Username, request.Password)
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "account registered succesfully",
			"token":   token,
		},
	)
}

func LoginUser(ctx *gin.Context) {
	var request requests.LoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		panic(err)
	}
	isUser := middleware.VerifyUser(request.Username, request.Password)
	if !isUser {
		ctx.SecureJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "dont hack this account",
			},
		)
		return
	}
	token, err := controllers.GenToken(request.Username)
	if err != nil {
		panic(err)
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "account logged in",
			"token":   token,
		},
	)
}
