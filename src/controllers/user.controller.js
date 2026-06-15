const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { errorResponse, sucessResponse } = require('../utils/responseFormatter');
const { UserService } = require('../services');

const signUser = async (req, res) => {
  try {
    const response = await UserService.signup({
      email: req.body.email,
      password: req.body.password,
    });
    LoggerConfig.info('Successfully recieved user Object from client');
    return res.status(StatusCodes.CREATED).json({
      ...sucessResponse,
      message: 'Successfully created a user object and send to database',
      data: response,
    });
  } catch (error) {
    LoggerConfig.error(`Error while creating booking: ${error.message}`);

    return res
      .status(error.statusCode || StatusCodes.INTERNAL_SERVER_ERROR)
      .json({
        ...errorResponse,
        message: error.message || 'Something went wrong',
        error: error,
      });
  }
};
module.exports = {
  signUser,
};
