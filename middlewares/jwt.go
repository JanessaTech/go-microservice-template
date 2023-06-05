package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func VerifyJwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("--------------- enter VerifyJwtToken -")
		ctx.Next()
	}
}
