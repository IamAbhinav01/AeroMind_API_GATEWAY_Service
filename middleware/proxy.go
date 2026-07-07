package middleware

import (
	config "AeromindGO/config/env"
	reverseproxy "AeromindGO/utils/reverse-proxy"
	"net/http"
)

var FlightsProxy http.Handler = reverseproxy.CreateProxyMiddleware(
	config.GetString("FLIGHTS_SERVICE_URL","FLIGHTS_SERVICE_URL"),"/api/v1/flights",
)

var BookingsProxy http.Handler = reverseproxy.CreateProxyMiddleware(
	config.GetString("BOOKING_SERVICE_URL","BOOKING_SERVICE_URL"),"/api/v1/booking",
)

var AIProxy http.Handler = reverseproxy.CreateProxyMiddleware(
	config.GetString("AI_SERVICE_URL","AI_SERVICE_URL"),"/api/v1/ai",
)