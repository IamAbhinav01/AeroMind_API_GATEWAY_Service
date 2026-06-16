const bcrypt = require('bcrypt');
const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository } = require('../repositories');
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
    console.log('the output of is_Valid_..', is_Valid_Password);
    return user;
  } catch (error) {
    if (error instanceof ErrorHandler) {
      throw error;
    }
    throw new ErrorHandler(
      'Error occured while signing in user',
      StatusCodes.BAD_REQUEST
    );
  }
};
module.exports = {
  signup,
  signin,
};
