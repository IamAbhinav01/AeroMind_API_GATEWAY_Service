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
const { RoleService } = require('../services');

const validatingAuthenticationRequest = async (req, res, next) => {
  const requiredFields = [Email, Password];
  for (const field of requiredFields) {
    let value = req.body ? req.body[field] : undefined;
    if (typeof value === 'string') value = value.trim();
    if (!value) {
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
    req.body[field] = value;
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

const isAdmin = async (req, res, next) => {
  try {
    const output = await RoleService.isAdmin(req.user.id);
    if (!output) {
      const responsePayload = {
        ...errorResponse,
        message: 'User is not authorised for this role',
        error: new ErrorHandler(message, StatusCodes.BAD_REQUEST),
      };
      return res.status(StatusCodes.BAD_REQUEST).json(responsePayload);
    }
    next();
  } catch (error) {
    const message = `Failed to verify the Admin Role`;
    const responsePayload = {
      ...errorResponse,
      message,
      error: new ErrorHandler(message, StatusCodes.BAD_REQUEST),
    };
    LoggerConfig.error(
      `Failed to verify the Admin role, ERROR Name: ${responsePayload.error.name}, ERROR Message: ${responsePayload.error.message}`
    );
    return res.status(StatusCodes.BAD_REQUEST).json(responsePayload);
  }
};
module.exports = {
  validatingAuthenticationRequest,
  AuthorizationBearerRequest,
  isAdmin,
};
