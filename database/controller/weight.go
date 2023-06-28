package controller

import (
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	timeUtils "github.com/kilianp07/MuscleApp/utils/time"
)

func (c *Controller) GetLatestWeight(userId uint) (*weightModel.Public, error) {
	var weight weightModel.Weight
	if err := c.db.Where("user_id = ?", userId).Last(&weight).Error; err != nil {
		return nil, err
	}
	return weightModel.ModelToPublic(&weight), nil
}

func (c *Controller) GetWeights(userId uint) ([]*weightModel.Public, error) {
	var (
		data    []*weightModel.Weight
		weights []*weightModel.Public
	)
	err := c.db.Where("user_id = ?", userId).Find(&data).Error

	for _, weight := range data {
		weights = append(weights, weightModel.ModelToPublic(weight))
	}

	if err != nil {
		return weights, err
	}
	return weights, nil
}

func (c *Controller) CreateWeight(weight *weightModel.Weight) error {
	return c.db.Create(&weight).Error
}

func (c *Controller) GetInitialWeight(userId uint) (*weightModel.Weight, error) {
	var weight weightModel.Weight
	if err := c.db.Where("user_id = ?", userId).First(&weight).Error; err != nil {
		return nil, err
	}
	return &weight, nil
}

func (c *Controller) DeleteWeight(id uint) error {
	return c.db.Delete(&weightModel.Weight{}, id).Error
}

func (c *Controller) DeleteWeightByDate(userId uint, date int) error {
	return c.db.Where("user_id = ? AND date = ?", userId, timeUtils.TimestampToTime(int64(date))).Delete(&weightModel.Weight{}).Error
}

func (c *Controller) UpdateWeight(weight *weightModel.Weight) error {
	return c.db.Save(&weight).Error
}

func (c *Controller) UpdateWeightByDate(userId uint, weight *weightModel.Public) error {
	return c.db.Model(&weightModel.Weight{}).Where("user_id = ? AND date = ?", userId, timeUtils.TimestampToTime(weight.Date)).Updates(weightModel.PublicToModel(weight, weight.Date, userId)).Error
}

func (c *Controller) GetWeightsBetween(userId uint, start int, end int) ([]*weightModel.Public, error) {
	var (
		data    []*weightModel.Weight
		weights []*weightModel.Public
	)
	err := c.db.Where("user_id = ? AND date BETWEEN ? AND ?", userId, timeUtils.TimestampToTime(int64(start)), timeUtils.TimestampToTime(int64(end))).Find(&data).Error

	for _, weight := range data {
		weights = append(weights, weightModel.ModelToPublic(weight))
	}

	if err != nil {
		return weights, err
	}
	return weights, nil
}
