package middlewares

import (
	"net/http"
	"time"
	"golang.org/x/time/rate"

)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 5) // 5 requests per second with a burst of 10

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}	
		next.ServeHTTP(w, r)
	})
}	