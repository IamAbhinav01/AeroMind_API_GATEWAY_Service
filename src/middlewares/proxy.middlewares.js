const { createProxyMiddleware } = require('http-proxy-middleware');
const {
  FLIGHTS_SERVICE_URL,
  BOOKING_SERVICE_URL,
} = require('../config/server.config');

const flightsProxy = createProxyMiddleware({
  target: FLIGHTS_SERVICE_URL,
  changeOrigin: true,
  proxyTimeout: 30000,
  timeout: 30000,
  // Express strips '/flights' from the path before the proxy sees it.
  // pathRewrite puts it back so backend receives /api/v1/flights/*
  pathRewrite: (path) => `/api/v1/flights${path}`,
  on: {
    error: (err, req, res) => {
      res.status(502).json({ error: 'Flights service unavailable', detail: err.message });
    },
  },
});

const bookingProxy = createProxyMiddleware({
  target: BOOKING_SERVICE_URL,
  changeOrigin: true,
  proxyTimeout: 30000,
  timeout: 30000,
  // Express strips '/bookings' from the path before the proxy sees it.
  // pathRewrite puts it back so backend receives /api/v1/booking/*
  pathRewrite: (path) => `/api/v1/booking${path}`,
  on: {
    error: (err, req, res) => {
      res.status(502).json({ error: 'Booking service unavailable', detail: err.message });
    },
  },
});

module.exports = {
  flightsProxy,
  bookingProxy,
};
