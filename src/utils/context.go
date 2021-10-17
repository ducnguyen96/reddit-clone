package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	GinContextKey    = "GinContext"
	AuthorizationKey = "Authorization"
	RequestTimeout   = 10
)

func GetAuthToken(ctx context.Context) string {
	headers := GetRequestHeaderFromContext(ctx)
	if auth, ok := headers[AuthorizationKey]; ok && len(auth) > 0 {
		return auth[0]
	}
	return ""
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetRequestHeaderFromContext(ctx context.Context) http.Header {
	ginCtx := GetGinContext(ctx)
	return ginCtx.Request.Header
}

func GetGinContext(ctx context.Context) *gin.Context {
	return ctx.Value(GinContextKey).(*gin.Context)
}