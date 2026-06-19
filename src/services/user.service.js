const { StatusCodes } = require('http-status-codes');
const { LoggerConfig, ServerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository } = require('../repositories');
const {
  checkPassword,
  jwtToken,
  verifyjwtToken,
} = require('../auth/logic.auth');
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
    console.log('JWT TOKEN: ', jwt_token);
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
      LoggerConfig.error('Token not found .');
      throw new ErrorHandler(
        'Token Not found , check again',
        StatusCodes.BAD_REQUEST
      );
    }
    const output = await verifyjwtToken(currentToken);
    const userDetails = await userRepo.get(output.id);
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
