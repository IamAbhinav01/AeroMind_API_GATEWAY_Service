const express = require('express');
const { UserController } = require('../../controllers');
const { Authentication, RateLimiter } = require('../../middlewares');

const router = express.Router();

router.post(
  '/signup',
  RateLimiter.strictLimiter,
  Authentication.validatingAuthenticationRequest,
  UserController.signUpUser
);
router.post(
  '/signin',
  RateLimiter.strictLimiter,
  Authentication.validatingAuthenticationRequest,
  UserController.signInUser
);

module.exports = router;
