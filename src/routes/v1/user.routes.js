const express = require('express');
const { UserController } = require('../../controllers');
const router = express.Router();

router.post('/signup', UserController.signUpUser);
router.post('/signin', UserController.signInUser);

module.exports = router;
