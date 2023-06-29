package sessionModel

import (
	sessionexerciseModel "github.com/kilianp07/MuscleApp/models/sessionExercise"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID          uint                                   `json:"id" gorm:"primary_key"`
	UserID      uint                                   `json:"user_id"`
	Exercises   []sessionexerciseModel.SessionExercise `json:"exercises"`
	Title       string                                 `json:"title"`
	Description string                                 `json:"description"`
}

type Public struct {
	ID          uint                          `json:"id"`
	Exercises   []sessionexerciseModel.Public `json:"exercises"`
	Title       string                        `json:"title"`
	Description string                        `json:"description"`
}

type Create struct {
	Exercises   []sessionexerciseModel.Create `json:"exercises" binding:"required"`
	Title       string                        `json:"title" binding:"required"`
	Description string                        `json:"description" binding:"required"`
}

func ModelToPublic(session *Session) Public {

	exercises := make([]sessionexerciseModel.Public, len(session.Exercises))
	for i := 0; i < len(session.Exercises); i++ {
		exercises[i] = sessionexerciseModel.ModelToPublic(session.Exercises[i])
	}

	return Public{
		ID:          session.ID,
		Exercises:   exercises,
		Title:       session.Title,
		Description: session.Description,
	}
}

func CreateToModel(session Create, userId uint) Session {
	exercises := make([]sessionexerciseModel.SessionExercise, len(session.Exercises))
	for i := 0; i < len(session.Exercises); i++ {
		exercises[i] = sessionexerciseModel.CreateToModel(session.Exercises[i])
	}
	return Session{
		Exercises:   exercises,
		Title:       session.Title,
		Description: session.Description,
		UserID:      userId,
	}
}

func PublicToModel(session Public) Session {
	exercises := make([]sessionexerciseModel.SessionExercise, len(session.Exercises))
	for i := 0; i < len(session.Exercises); i++ {
		exercises[i] = sessionexerciseModel.PublicToModel(session.Exercises[i])
	}
	return Session{
		ID:          session.ID,
		Exercises:   exercises,
		Title:       session.Title,
		Description: session.Description,
	}
}
