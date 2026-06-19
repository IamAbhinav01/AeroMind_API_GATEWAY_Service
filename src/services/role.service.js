const { StatusCodes } = require('http-status-codes');
const { LoggerConfig } = require('../config');
const { ErrorHandler } = require('../errors');
const { UserRepository, RoleRepository } = require('../repositories');
const { ADMIN } = require('../utils/common/roles.constants');

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
    const roleName = data.role ? data.role.toString().trim() : '';
    if (!roleName) {
      throw new ErrorHandler('Role not provided', StatusCodes.BAD_REQUEST);
    }
    const roleDetails = await roleRepository.getRoleDetails(roleName);
    if (!Array.isArray(roleDetails) || roleDetails.length === 0) {
      throw new ErrorHandler(`No role found for ${roleName}`, StatusCodes.BAD_REQUEST);
    }
    const roleInstance = roleDetails[0];
    LoggerConfig.info(`Role details found for role=${roleName}`);
    LoggerConfig.info(`User details found for id=${data.id}`);
    await userDetails.addRole(roleInstance);
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

const isAdmin = async (id) => {
  try {
    const userDetails = await userRepository.get(id);
    if (!userDetails) {
      throw new ErrorHandler(
        'No user found for this credential provided',
        StatusCodes.BAD_REQUEST
      );
    }
    const isAdminRole = await roleRepository.getRoleDetails(ADMIN);
    if (!Array.isArray(isAdminRole) || isAdminRole.length === 0) {
      throw new ErrorHandler(`No Admin role found`, StatusCodes.BAD_REQUEST);
    }
    const adminRoleInstance = isAdminRole[0];
    return userDetails.hasRole(adminRoleInstance);
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
  isAdmin,
};
