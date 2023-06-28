package controller

import objectiveModel "github.com/kilianp07/MuscleApp/models/objective"

func (c *Controller) CreateObjective(objective *objectiveModel.Objective) error {
	return c.db.Create(&objective).Error
}

func (c *Controller) GetObjectiveByUserID(userId uint) (*objectiveModel.Public, error) {
	var objective objectiveModel.Objective
	if err := c.db.Where("user_id = ?", userId).First(&objective).Error; err != nil {
		return nil, err
	}
	return objectiveModel.ModelToPublic(&objective), nil
}

func (c *Controller) GetObjectiveByID(id int) (*objectiveModel.Objective, error) {
	var objective objectiveModel.Objective
	if err := c.db.Where("ID = ?", id).First(&objective).Error; err != nil {
		return nil, err
	}
	return &objective, nil
}

func (c *Controller) UpdateObjective(objective *objectiveModel.Objective, id int) error {
	// Get the existing Objective record from the database
	var existingObjective objectiveModel.Objective
	if err := c.db.First(&existingObjective, id).Error; err != nil {
		return err
	}

	// Update the fields with the new values
	existingObjective.Title = objective.Title
	existingObjective.Weight = objective.Weight
	existingObjective.Description = objective.Description
	existingObjective.Date = objective.Date

	// Perform the update operation
	if err := c.db.Save(&existingObjective).Error; err != nil {
		return err
	}

	return nil
}

func (c *Controller) DeleteObjective(objective *objectiveModel.Objective) error {
	return c.db.Delete(&objective).Where("ID = ?", objective.ID).Error
}

func (c *Controller) GetAllObjectives(userId uint) ([]objectiveModel.Public, error) {
	var (
		data       []objectiveModel.Objective
		objectives []objectiveModel.Public
	)
	if err := c.db.Where("user_id = ?", userId).Find(&data).Error; err != nil {
		return nil, err
	}

	for _, objective := range data {
		objectives = append(objectives, *objectiveModel.ModelToPublic(&objective))
	}

	return objectives, nil
}
