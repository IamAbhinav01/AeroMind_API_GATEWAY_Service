const express = require('express');
const { RoleController } = require('../../controllers');
const router = express.Router();

router.post('/assign', RoleController.AssignRole);

module.exports = router;
