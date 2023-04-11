'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class WeightGoal extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      WeightGoal.hasOne(models.User)
    }
  }
  WeightGoal.init({
    value: DataTypes.FLOAT,
    timestamp: DataTypes.DATE,
    isDeleted: DataTypes.BOOLEAN,
  }, {
    sequelize,
    modelName: 'WeightGoal',
  });
  return WeightGoal;
};