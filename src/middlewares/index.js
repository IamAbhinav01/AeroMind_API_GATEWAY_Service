module.exports = {
  Authentication: require('./authentication.midlewares'),
  RateLimiter: require('./ratelimiter.middleware'),
  ProxyMiddleware: require('./proxy.middlewares'),
};
