const express = require('express');
const { RoleController } = require('../../controllers');
const {
  AuthorizationBearerRequest,
  isAdmin,
} = require('../../middlewares/authentication.midlewares');
const router = express.Router();

router.post(
  '/assign',
  AuthorizationBearerRequest,
  isAdmin,
  RoleController.AssignRole
);

module.exports = router;
