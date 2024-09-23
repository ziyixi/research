package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware is a middleware that limits the number of requests per minute
func RateLimitMiddleware() gin.HandlerFunc {
	// Declare variables inside the closure
	var mu sync.Mutex
	requestsCount := 0
	resetTime := time.Now().Add(1 * time.Minute)

	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()

		// Check if the time window has expired
		if time.Now().After(resetTime) {
			// Reset the counter and the time window
			requestsCount = 0
			resetTime = time.Now().Add(1 * time.Minute)
		}

		// Check the request count
		if requestsCount >= 2 {
			// Block the request if the limit is reached
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Too many requests. Please wait for the next minute.",
			})
			c.Abort()
			return
		}

		// Allow the request and increment the counter
		requestsCount++

		// Process the request
		c.Next()
	}
}
