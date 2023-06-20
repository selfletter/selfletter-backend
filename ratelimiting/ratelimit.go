package ratelimiting

import (
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RateLimitResponse struct {
	Error string `json:"error"`
}

func RateLimitKeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func RateLimitErrorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(http.StatusTooManyRequests, RateLimitResponse{Error: "try again in " + time.Until(info.ResetTime).Truncate(time.Second).String()})
}

func GetRateLimitHandler(interval uint) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Duration(interval) * time.Second,
		Limit: 1,
	})

	rateLimitHandler := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: RateLimitErrorHandler,
		KeyFunc:      RateLimitKeyFunc,
	})

	return rateLimitHandler
}
