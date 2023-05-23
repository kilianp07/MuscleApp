package controller

import (
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	timeUtils "github.com/kilianp07/MuscleApp/utils/time"
)

func (c *Controller) GetLatestWeight(userId uint) (*weightModel.Model, error) {
	var weight weightModel.Model
	if err := c.db.Where("user_id = ?", userId).Last(&weight).Error; err != nil {
		return &weight, err
	}
	return &weight, nil
}

func (c *Controller) GetWeights(userId uint) ([]*weightModel.Public, error) {
	var (
		data    []*weightModel.Model
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

func (c *Controller) CreateWeight(weight *weightModel.Model) error {
	if err := c.db.Create(&weight).Error; err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteWeight(id uint) error {
	if err := c.db.Delete(&weightModel.Model{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteWeightByDate(userId uint, date int) error {
	if err := c.db.Where("user_id = ? AND date = ?", userId, timeUtils.TimestampToTime(int64(date))).Delete(&weightModel.Model{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateWeight(weight *weightModel.Model) error {
	if err := c.db.Save(&weight).Error; err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateWeightByDate(userId uint, weight *weightModel.Public) error {
	if err := c.db.Model(&weightModel.Model{}).Where("user_id = ? AND date = ?", userId, timeUtils.TimestampToTime(weight.Date)).Updates(weightModel.PublicToModel(weight, weight.Date, userId)).Error; err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetWeightsBetween(userId uint, start int, end int) ([]*weightModel.Public, error) {
	var (
		data    []*weightModel.Model
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
