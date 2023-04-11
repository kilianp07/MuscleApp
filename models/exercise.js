'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class Exercise extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      Exercise.hasMany(models.Muscle)
      Exercise.hasMany(models.Tag)
    }
  }
  Exercise.init({
    title: DataTypes.STRING,
    descritpion: DataTypes.TEXT('long'),
    isDeleted: DataTypes.BOOLEAN
  }, {
    sequelize,
    modelName: 'Exercise',
  });
  return Exercise;
};