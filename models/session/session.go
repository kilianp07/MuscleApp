package sessionModel

import exerciseModel "github.com/kilianp07/MuscleApp/models/Exercise"

type Session struct {
	ID          uint                   `json:"id" gorm:"primaryKey"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Exercises   []exerciseModel.Public `json:"exercises" gorm:"many2many:session_exercises;"`
	UserID      uint                   `json:"user_id"`
}

type Public struct {
	ID          uint                   `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Exercises   []exerciseModel.Public `json:"exercises"`
}

func ModelToPublic(session *Session) *Public {
	return &Public{
		ID:          session.ID,
		Title:       session.Title,
		Description: session.Description,
		Exercises:   session.Exercises,
	}
}

func PublicToModel(public *Public, user_id uint) *Session {
	return &Session{
		ID:          public.ID,
		Title:       public.Title,
		Description: public.Description,
		Exercises:   public.Exercises,
		UserID:      user_id,
	}
}
