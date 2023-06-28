package controller

import exerciseModel "github.com/kilianp07/MuscleApp/models/Exercise"

func (c *Controller) CreateExercise(exercise *exerciseModel.Exercise) error {
	return c.db.Create(&exercise).Error
}

func (c *Controller) GetExerciseByID(id int) (*exerciseModel.Public, error) {
	var exercise exerciseModel.Exercise
	if err := c.db.Where("ID = ?", id).First(&exercise).Error; err != nil {
		return nil, err
	}
	return exerciseModel.ModelToPublic(&exercise), nil
}

func (c *Controller) GetSomeExercises(number int) ([]exerciseModel.Public, error) {
	var (
		data      []exerciseModel.Exercise
		exercises []exerciseModel.Public
	)
	if err := c.db.Limit(number).Find(&data).Error; err != nil {
		return nil, err
	}

	for _, exercise := range data {
		exercises = append(exercises, *exerciseModel.ModelToPublic(&exercise))
	}

	return exercises, nil
}

func (c *Controller) UpdateExercise(data *exerciseModel.Exercise) error {
	var exercise exerciseModel.Exercise
	if err := c.db.Where("ID = ?", data.ID).First(&exercise).Error; err != nil {
		return err
	}

	exercise = exerciseModel.Exercise{
		Title:       data.Title,
		Description: data.Description,
		Video:       data.Video,
		Difficulty:  data.Difficulty,
		Member:      data.Member,
		Type:        data.Type,
	}

	return c.db.Save(&exercise).Where("ID = ?", exercise.ID).Error
}

func (c *Controller) DeleteExercise(exercise *exerciseModel.Exercise) error {
	return c.db.Delete(&exercise).Where("ID = ?", exercise.ID).Error
}

func (c *Controller) GetLatestExercise() (*exerciseModel.Public, error) {
	var exercise exerciseModel.Exercise
	err := c.db.Order("id desc").First(&exercise).Error
	if err != nil {
		return nil, err
	}
	return exerciseModel.ModelToPublic(&exercise), nil
}
