package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/helper"
)

func VerifyJwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("--------------- enter VerifyJwtToken -")
		if err := helper.ValidateJWT(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "anthentication is failed"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
