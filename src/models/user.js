'use strict';
const { Model } = require('sequelize');
const bcrypt = require('bcrypt');
const serverConfig = require('../config/server.config');
module.exports = (sequelize, DataTypes) => {
  class User extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  }
  User.init(
    {
      email: {
        type: DataTypes.STRING,
        unique: true,
        allowNull: false,
        validate: { isEmail: true },
      },
      password: {
        type: DataTypes.STRING,
        allowNull: false,
        validate: { len: [6, 20] },
      },
    },

    {
      sequelize,
      modelName: 'User',

      hooks: {
        beforeCreate: async (user) => {
          if (!user.password) return;
          const salt = await bcrypt.genSaltSync(+serverConfig.SALT_ROUNDS, 'a');
          user.password = bcrypt.hashSync(user.password, salt);
        },
      },
    }
  );
  return User;
};
