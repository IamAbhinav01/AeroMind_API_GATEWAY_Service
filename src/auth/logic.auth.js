const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const { StatusCodes } = require('http-status-codes');
const { LoggerConfig, ServerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { JWT_EXPIRY, JWT_SECRET } = require('../config/server.config');

const checkPassword = async (inputPassword, backendPassword) => {
  try {
    const is_match = await bcrypt.compare(inputPassword, backendPassword);
    return is_match;
  } catch (error) {
    LoggerConfig.error('Error while checking the password');
    throw new ErrorHandler(
      'Error while checking the password ',
      StatusCodes.BAD_GATEWAY
    );
  }
};

const jwtToken = async (incommingRequest) => {
  try {
    return jwt.sign(incommingRequest, JWT_SECRET, { expiresIn: JWT_EXPIRY });
  } catch (error) {
    LoggerConfig.error('Error while craeting the jwt token');
    throw new ErrorHandler(
      'Error while craeting the jwt token ',
      StatusCodes.BAD_GATEWAY
    );
  }
};

const verifyjwtToken = async (currentToken) => {
  try {
    return jwt.verify(currentToken, JWT_SECRET);
  } catch (error) {
    LoggerConfig.error('Error while verifying  the jwt token');
    throw new ErrorHandler(
      'Error while verifying the jwt token ',
      StatusCodes.BAD_GATEWAY
    );
  }
};

module.exports = {
  checkPassword,
  jwtToken,
  verifyjwtToken,
};
