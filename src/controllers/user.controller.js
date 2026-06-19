const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { errorResponse, sucessResponse } = require('../utils/responseFormatter');
const { UserService } = require('../services');

const signUpUser = async (req, res) => {
  try {
    const email = req.body.email ? req.body.email.toString().trim() : '';
    const password = req.body.password ? req.body.password.toString() : '';
    const response = await UserService.signup({
      email,
      password,
    });
    const userObject = response.toJSON();
    delete userObject.password;
    LoggerConfig.info('Successfully recieved user Object from client');
    return res.status(StatusCodes.CREATED).json({
      ...sucessResponse,
      message: 'Successfully created a user object and send to database',
      data: userObject,
    });
  } catch (error) {
    LoggerConfig.error(`Error while creating user: ${error.message}`);

    return res
      .status(error.statusCode || StatusCodes.INTERNAL_SERVER_ERROR)
      .json({
        ...errorResponse,
        message: error.message || 'Something went wrong',
        error: error,
      });
  }
};

const signInUser = async (req, res) => {
  try {
    const email = req.body.email ? req.body.email.toString().trim() : '';
    const password = req.body.password ? req.body.password.toString() : '';
    const response = await UserService.signin({
      email,
      password,
    });

    return res.status(StatusCodes.ACCEPTED).json({
      ...sucessResponse,
      message: 'Successfully singIn the user',
      data: response,
    });
  } catch (error) {
    LoggerConfig.error(`Error while signIn the user: ${error.message}`);

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
  signUpUser,
  signInUser,
};
