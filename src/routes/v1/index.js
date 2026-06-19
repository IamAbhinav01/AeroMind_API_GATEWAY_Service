const express = require('express');
const { healthController } = require('../../controllers/');
const userRoutes = require('./user.routes');
const roleRoutes = require('./role.routes');
const { RateLimiter } = require('../../middlewares');
const {
  flightsProxy,
  bookingProxy,
} = require('../../middlewares/proxy.middlewares');
const { checkAuthentication } = require('../../services/user.service');
const {
  AuthorizationBearerRequest,
} = require('../../middlewares/authentication.midlewares');

const router = express.Router();

router.use('/flights', RateLimiter.standardLimiter, flightsProxy);
router.use(
  '/bookings',
  RateLimiter.standardLimiter,
  AuthorizationBearerRequest,
  bookingProxy
);

router.use('/healthy', healthController.health);
router.use('/user', userRoutes);
router.use('/user/role', roleRoutes);
module.exports = router;
