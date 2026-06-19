const express = require('express');
const { healthController } = require('../../controllers/');
const userRoutes = require('./user.routes');
const roleRoutes = require('./role.routes');

const router = express.Router();

router.use('/healthy', healthController.health);
router.use('/user', userRoutes);
router.use('/user/role', roleRoutes);
module.exports = router;
