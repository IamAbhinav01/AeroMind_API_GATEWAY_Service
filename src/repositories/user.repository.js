const { CrudRepository } = require('./crudOperations.repository');
const { User } = require('../models');
const { LoggerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { StatusCodes } = require('http-status-codes');
class UserRepository extends CrudRepository {
  constructor() {
    super(User);
  }
  async getUserDetails(email) {
    try {
      const user = await User.findOne({
        where: {
          email: email,
        },
      });
      if (!user) {
        LoggerConfig.error('No user found for this email id ');
        throw new ErrorHandler(
          'Error while finding the user with given email ID',
          StatusCodes.BAD_GATEWAY
        );
      }
      return user;
    } catch (error) {
      if (error instanceof ErrorHandler) {
        throw error;
      }
      LoggerConfig.error(
        `Error occured while fetching the user with email ${email} ERROR:${error}`
      );
      throw new ErrorHandler(
        `Error occured while fetching the user with email ${email} ERROR:${error}`,
        StatusCodes.INTERNAL_SERVER_ERROR
      );
    }
  }
}

module.exports = UserRepository;
