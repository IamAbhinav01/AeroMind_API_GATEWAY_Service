'use strict';
const { Model } = require('sequelize');
const { ADMIN, USER, AIRLINE } = require('../utils/common/roles.constants');

module.exports = (sequelize, DataTypes) => {
  class Role extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  }
  Role.init(
    {
      name: {
        type: DataTypes.ENUM({ values: [ADMIN, USER, AIRLINE] }),
        allowNull: false,
        defaultValue: USER,
      },
    },
    {
      sequelize,
      modelName: 'Role',
    }
  );
  return Role;
};
