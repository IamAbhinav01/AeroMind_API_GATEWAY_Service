const { CrudRepository } = require('./crudOperations.repository');
const { Role } = require('../models');

class RoleRepository extends CrudRepository {
  constructor() {
    super(Role);
  }
  async getRoleDetails(name) {
    const roleDetails = await Role.findAll({
      where: { name: name },
    });
    return roleDetails;
  }
}
module.exports = RoleRepository;
