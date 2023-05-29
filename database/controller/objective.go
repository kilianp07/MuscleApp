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

func (c *Controller) UpdateObjective(objective *objectiveModel.Objective) error {
	return c.db.Save(&objective).Error
}

func (c *Controller) DeleteObjective(objective *objectiveModel.Objective) error {
	return c.db.Delete(&objective).Error
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
