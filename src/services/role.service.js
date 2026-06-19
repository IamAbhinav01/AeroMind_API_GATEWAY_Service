const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository, RoleRepository } = require('../repositories');

const userRepository = new UserRepository();
const roleRepository = new RoleRepository();

const AssignRoleToUser = async (data) => {
  try {
    const userDetails = await userRepository.get(data.id);
    if (!userDetails) {
      throw new ErrorHandler(
        'No user found for this credential provided',
        StatusCodes.BAD_REQUEST
      );
    }
    const roleDetails = await roleRepository.getRoleDetails(data.role);
    if (!Array.isArray(roleDetails) || roleDetails.length === 0) {
      throw new ErrorHandler(
        `No role found for ${data.role}`,
        StatusCodes.BAD_REQUEST
      );
    }
    LoggerConfig.info(`Role details found for role=${data.role}`);
    LoggerConfig.info(`User details found for id=${data.id}`);
    return userDetails;
  } catch (error) {
    if (error instanceof ErrorHandler) {
      throw error;
    }
    LoggerConfig.error('error while Assignning  the user the role');
    throw new ErrorHandler(
      'Error occured while Permssion transfer to user',
      StatusCodes.BAD_REQUEST
    );
  }
};
module.exports = {
  AssignRoleToUser,
};
