const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { RoleService } = require('../services');
const { sucessResponse, errorResponse } = require('../utils/responseFormatter');

const AssignRole = async (req, res) => {
  try {
    const response = await RoleService.AssignRoleToUser({
      id: req.body.id,
      role: req.body.role,
    });

    return res.status(StatusCodes.ACCEPTED).json({
      ...sucessResponse,
      message: 'Successfully Assigned the user a role',
      data: response,
    });
  } catch (error) {
    LoggerConfig.error(
      `Error while  Assigned the user a role: ${error.message}`
    );

    return res
      .status(error.statusCode || StatusCodes.INTERNAL_SERVER_ERROR)
      .json({
        ...errorResponse,
        message: error.message || 'Something went wrong',
        error: error,
      });
  }
};
module.exports = {
  AssignRole,
};
