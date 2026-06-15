const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository } = require('../repositories');
const userRepo = new UserRepository();
const signup = async function (data) {
  try {
    const userData = userRepo.create(data);
    LoggerConfig.info('successfully user and stored in database ');
    if (!userData) {
      LoggerConfig.error('Error while registering user to database');
      throw new ErrorHandler(
        'Error while registering user to database',
        StatusCodes.INTERNAL_SERVER_ERROR
      );
    }
    return userData;
  } catch (error) {
    console.log(error);
    if (error instanceof ErrorHandler) {
      throw error;
    }
    throw new ErrorHandler(
      'Cannot store user to database.',
      StatusCodes.INTERNAL_SERVER_ERROR
    );
  }
};
module.exports = {
  signup,
};
