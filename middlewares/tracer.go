package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hi-supergirl/go-microservice-template/logging"
)

const (
	XRequestIdKey = string("X-Request-ID") // request id header key
)

func RequestTraceMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get(XRequestIdKey)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx.Request = ctx.Request.WithContext(contextWithRequestId(ctx, requestId))
		logger := logging.GetLogger(true).Sugar().With("requestId", requestId)
		ctx.Request = ctx.Request.WithContext(logging.WithLogger(ctx, logger))
		ctx.Writer.Header().Set(XRequestIdKey, requestId)
	}
}

func contextWithRequestId(ctx context.Context, requestId string) context.Context {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}
	return context.WithValue(ctx, "requestId", requestId)
}
