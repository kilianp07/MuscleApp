'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class BodyMember extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
      BodyMember.hasMany(models.Muscle)
      BodyMember.hasMany(models.Exercise)
    }
  }
  BodyMember.init({
    name: DataTypes.STRING,
    isDeleted:  DataTypes.BOOLEAN,
  }, {
    sequelize,
    modelName: 'BodyMember',
  });
  return BodyMember;
};