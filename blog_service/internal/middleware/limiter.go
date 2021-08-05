package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/errcode"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/limit"
)

func RateLimiter(l limit.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(bucket.Capacity())
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
