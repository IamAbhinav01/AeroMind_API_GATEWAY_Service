package middleware

import (
	"AeromindGO/utils/limiter"
	"net/http"
)

type RateLimiterOptions struct {
	Label     string
	Capacity  int
	Refill    int
	Timeout   int
	Cost      int
	Whitelist []string
	KeyFunc   func(r *http.Request) string
	OnDenied  func(w http.ResponseWriter, r *http.Request, result *limiter.BucketResponse)
}

func createRateLimiter(
	service RedisBucketService,
	options RateLimiterOptions,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			next.ServeHTTP(w, r)

		})
	}
}