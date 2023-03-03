'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class ExerciseProgram extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      ExerciseProgram.hasOne(models.User)
      ExerciseProgram.hasMany(models.Exercise)
      ExerciseProgram.hasMany(models.Tag)
    }
  }
  ExerciseProgram.init({
    progression: DataTypes.INTEGER,
    title: DataTypes.STRING
  }, {
    sequelize,
    modelName: 'ExerciseProgram',
  });
  return ExerciseProgram;
};