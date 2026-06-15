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
module.exports = {
  signup,
};
