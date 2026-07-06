package middleware

import (
	"AeromindGO/utils/limiter"
	"math"
	"net/http"
	"strconv"
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

func CreateRateLimiter(
	service limiter.RedisBucketService,
	options RateLimiterOptions,
) func(http.Handler) http.Handler {
	if options.KeyFunc == nil {
		options.KeyFunc = func(r *http.Request) string {
			return r.RemoteAddr
		}
	}

	if options.Cost == 0 {
		options.Cost = 1
	}

	whitelist := make(map[string]struct{})

	for _, ip := range options.Whitelist {
		whitelist[ip] = struct{}{}
	}

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			identifier := options.KeyFunc(r)

			if _, ok := whitelist[identifier]; ok {
				w.Header().Set("X-RateLimit-Allowed", "true")
				next.ServeHTTP(w, r)
				return
			}

			result, err := service.IsRequestAllowed(
				identifier,
				&limiter.BucketOptions{
					Capacity: options.Capacity,
					Refill:   options.Refill,
					Timeout:  options.Timeout,
					Cost:     options.Cost,
				},
			)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set(
				"X-RateLimit-Limit",
				strconv.Itoa(result.Capacity),
			)

			w.Header().Set(
				"X-RateLimit-Remaining",
				strconv.Itoa(result.Remaining),
			)

			if !result.Allowed {

				retry := int(math.Ceil(
					float64(result.RetryAfterMs) / 1000,
				))

				w.Header().Set(
					"Retry-After",
					strconv.Itoa(retry),
				)

				if options.OnDenied != nil {
					options.OnDenied(w, r, result)
					return
				}

				http.Error(
					w,
					"Too Many Requests",
					http.StatusTooManyRequests,
				)

				return
			}
			
			next.ServeHTTP(w, r)

		})
	}
}

