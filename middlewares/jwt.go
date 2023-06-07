package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/helper"
	"github.com/hi-supergirl/go-microservice-template/logging"
)

func JwtTokenMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := logging.FromContext(ctx)
		logger.Debugw("JwtTokenMiddleWare")
		if err := helper.ValidateJWT(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "anthentication is failed"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
