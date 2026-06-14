'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    /**
     * Add seed commands here.
     *
     * Example:
     * await queryInterface.bulkInsert('People', [{
     *   name: 'John Doe',
     *   isBetaMember: false
     * }], {});
     */
    await queryInterface.bulkInsert('Users', [
      {
        email: 'john.doe@example.com',
        password: 'password123',
        createdAt: new Date(),
        updatedAt: new Date(),
      },
      {
        email: 'jane.smith@example.com',
        password: 'password456',
        createdAt: new Date(),
        updatedAt: new Date(),
      },
      {
        email: 'bob.johnson@example.com',
        password: 'password789',
        createdAt: new Date(),
        updatedAt: new Date(),
      },
      {
        email: 'ogg@example.com',
        password: 'password000',
        createdAt: new Date(),
        updatedAt: new Date(),
      },
    ]);
  },

  async down(queryInterface, Sequelize) {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
    await queryInterface.bulkDelete('Users', null, {});
  },
};
