package limit

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterInterface {
	l := &Limiter{
		limiterBuckets: make(map[string]*ratelimit.Bucket),
	}
	return MethodLimiter{
		Limiter: l,
	}
}

// Key 返回请求uri的路径
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

// GetBucket 获取限流桶
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// AddBuckets 添加限流桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}

	return l
}
