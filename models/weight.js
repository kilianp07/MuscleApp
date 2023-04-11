'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class Weight extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      // Weight has one User
      Weight.belongsTo(models.User)
    }
  }
  Weight.init({
    value: DataTypes.FLOAT,
    timestamp: DataTypes.DATE,
    isDeleted: DataTypes.BOOLEAN,
  }, {
    sequelize,
    modelName: 'Weight',
  });

  function incomingCorrectlyFilled(weight) {
    return weight.value && weight.timestamp && weight.userId;
  }

  return Weight;
};