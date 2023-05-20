package controller

import (
	"time"

	weightModel "github.com/kilianp07/MuscleApp/models/weight"
)

func (c *Controller) GetLatestWeight(userId uint) (*weightModel.Model, error) {
	var weight weightModel.Model
	if err := c.db.Where("user_id = ?", userId).Last(&weight).Error; err != nil {
		return &weight, err
	}
	return &weight, nil
}

func (c *Controller) GetWeights(userId uint) ([]*weightModel.Model, error) {
	var weights []*weightModel.Model
	if err := c.db.Where("user_id = ?", userId).Find(&weights).Error; err != nil {
		return weights, err
	}
	return weights, nil
}

func (c *Controller) GetWeightsBetweenDates(userId uint, startDate *time.Time, endDate *time.Time) ([]*weightModel.Model, error) {
	var weights []*weightModel.Model
	if err := c.db.Where("user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate).Find(&weights).Error; err != nil {
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
