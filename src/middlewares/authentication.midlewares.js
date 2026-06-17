const { StatusCodes } = require('http-status-codes');
const { ErrorHandler } = require('../errors');
const { errorResponse } = require('../utils/responseFormatter');
const { LoggerConfig } = require('../config');
const { Email, Password } = require('../utils/common/middleware_constants');

const validatingAuthenticationRequest = async (req, res, next) => {
  const requiredFields = [Email, Password];
  for (const field of requiredFields) {
    if (!req.body[field]) {
      const message = `${field} is not defined`;
      const responsePayload = {
        ...errorResponse,
        message,
        error: new ErrorHandler(message, StatusCodes.BAD_REQUEST),
      };
      LoggerConfig.error(
        ` ${field} not defined, ERROR Name: ${responsePayload.error.name}, ERROR Message: ${responsePayload.error.message}`
      );
      return res.status(StatusCodes.BAD_REQUEST).json(responsePayload);
    }
  }
  next();
};

module.exports = {
  validatingAuthenticationRequest,
};
