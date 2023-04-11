'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class User extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      User.hasMany(models.Weight)
      User.hasMany(models.WeightGoal)
      User.hasMany(models.ExerciseProgram)
    }
  }
  User.init({
    firstName: DataTypes.STRING,
    lastName: DataTypes.STRING,
    email: DataTypes.STRING,
    uid: DataTypes.STRING,
    password: DataTypes.STRING,
    isDeleted: DataTypes.BOOLEAN
    
  }, {
    sequelize,
    modelName: 'User',
  });
  return User;
};