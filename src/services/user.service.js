const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const { StatusCodes } = require('http-status-codes');
const { LoggerConfig, ServerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository } = require('../repositories');
const { JWT_EXPIRY, JWT_SECRET } = require('../config/server.config');
const userRepo = new UserRepository();
const signup = async function (data) {
  try {
    const userData = await userRepo.create(data);

    if (!userData) {
      LoggerConfig.error('Error while registering user to database');
      throw new ErrorHandler(
        'Error while registering user to database',
        StatusCodes.INTERNAL_SERVER_ERROR
      );
    }
    LoggerConfig.info('successfully user and stored in database ');
    return userData;
  } catch (error) {
    console.log(error);
    if (error instanceof ErrorHandler) {
      throw error;
    }
    if (error.name === 'SequelizeUniqueConstraintError') {
      throw new ErrorHandler(error.name, StatusCodes.INTERNAL_SERVER_ERROR);
    }
    throw new ErrorHandler(
      'Cannot store user to database.',
      StatusCodes.INTERNAL_SERVER_ERROR
    );
  }
};

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

const signin = async function (data) {
  try {
    const user = await userRepo.getUserDetails(data.email);
    const is_Valid_Password = await checkPassword(data.password, user.password);
    if (!is_Valid_Password) {
      throw new ErrorHandler(
        'Invalid email or password, check again and try',
        StatusCodes.BAD_REQUEST
      );
    }
    const jwt_token = await jwtToken({ id: user.id, email: user.email });
    LoggerConfig.info('successfully signed in the user');
    return jwt_token;
  } catch (error) {
    if (error instanceof ErrorHandler) {
      throw error;
    }
    LoggerConfig.error('error while signin the user');
    throw new ErrorHandler(
      'Error occured while signing in user',
      StatusCodes.BAD_REQUEST
    );
  }
};

const checkAuthentication = async (currentToken) => {
  try {
    if (!currentToken) {
      LoggerConfig('Token not found .');
      throw new ErrorHandler(
        'Token Not found , check again',
        StatusCodes.BAD_REQUEST
      );
    }
    const output = verifyjwtToken(currentToken);
    const userDetails = userRepo.get(output.id);
    if (!userDetails) {
      throw new ErrorHandler(
        'No user signature found in the jwt token',
        StatusCodes.CONFLICT
      );
    }
    return userDetails;
  } catch (error) {
    if (error instanceof ErrorHandler) {
      throw error;
    }
    throw new ErrorHandler(
      'Error occured while authenticating the user',
      StatusCodes.BAD_REQUEST
    );
  }
};

module.exports = {
  signup,
  signin,
  checkAuthentication,
};
