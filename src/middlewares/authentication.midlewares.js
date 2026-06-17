const { StatusCodes } = require('http-status-codes');
const { ErrorHandler } = require('../errors');
const { errorResponse } = require('../utils/responseFormatter');
const { LoggerConfig } = require('../config');
const {
  Email,
  Password,
  X_ACCESS_TOKEN,
} = require('../utils/common/middleware_constants');
const { checkAuthentication } = require('../services/user.service');

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

const AuthorizationBearerRequest = async (req, res, next) => {
  try {
    const AuthenticateUser = await checkAuthentication(
      req.headers[X_ACCESS_TOKEN]
    );
    req.user = AuthenticateUser;
    LoggerConfig.info(`Verified the bearer heads`);
    next();
  } catch (error) {
    const message = `Failed to verify the authenticity of the token`;
    const responsePayload = {
      ...errorResponse,
      message,
      error: new ErrorHandler(message, StatusCodes.BAD_REQUEST),
    };
    LoggerConfig.error(
      `Failed to verify the authenticity of the token, ERROR Name: ${responsePayload.error.name}, ERROR Message: ${responsePayload.error.message}`
    );
    return res.status(StatusCodes.BAD_REQUEST).json(responsePayload);
  }
};
module.exports = {
  validatingAuthenticationRequest,
  AuthorizationBearerRequest,
};
