const dotenv = require('dotenv').config();

module.exports = {
  // ── Server ─────────────────────────────────────────────────────────────────
  PORT: process.env.PORT,
  logger_level: process.env.logger_level,
  SALT_ROUNDS: process.env.SALT_ROUNDS,
  JWT_SECRET: process.env.JWT_SECRET,
  JWT_EXPIRY: process.env.JWT_EXPIRY,

  REDIS_HOST: process.env.REDIS_HOST || '127.0.0.1',
  REDIS_PORT: parseInt(process.env.REDIS_PORT, 10) || 6379,
  REDIS_TIMEOUT: parseInt(process.env.REDIS_TIMEOUT, 10) || 2000,

  BUCKET_CAPACITY: parseInt(process.env.BUCKET_CAPACITY, 10) || 100,
  BUCKET_REFILL:
    parseInt(process.env.BUCKET_REFILL || process.env.BUCKET_REFIL, 10) || 10,
  BUCKET_TIMEOUT: parseInt(process.env.BUCKET_TIMEOUT, 10) || 60000,
  FLIGHTS_SERVICE_URL: process.env.FLIGHTS_SERVICE_URL,
  BOOKING_SERVICE_URL: process.env.BOOKING_SERVICE_URL,
  FRONTEND_SERVICE_URL: process.env.FRONTEND_SERVICE_URL,
};
