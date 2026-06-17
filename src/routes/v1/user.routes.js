const express = require('express');
const { UserController } = require('../../controllers');
const { Authentication } = require('../../middlewares');
const router = express.Router();

router.post(
  '/signup',
  Authentication.validatingAuthenticationRequest,
  UserController.signUpUser
);
router.post(
  '/signin',
  Authentication.validatingAuthenticationRequest,

  UserController.signInUser
);

module.exports = router;
