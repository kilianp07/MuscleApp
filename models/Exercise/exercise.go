package exerciseModel

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       string `json:"video"`
	Difficulty  uint   `json:"difficulty"`
	Member      string `json:"member"`
	Type        string `json:"type"`
}

type Public struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       string `json:"video"`
	Difficulty  uint   `json:"difficulty"`
	Member      string `json:"member"`
	Type        string `json:"type"`
}

type Create struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Video       string `json:"video" binding:"required"`
	Difficulty  uint   `json:"difficulty" binding:"required"`
	Member      string `json:"member" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

func ModelToPublic(exercise *Exercise) *Public {
	return &Public{
		ID:          exercise.ID,
		Title:       exercise.Title,
		Description: exercise.Description,
		Video:       exercise.Video,
		Difficulty:  exercise.Difficulty,
		Member:      exercise.Member,
		Type:        exercise.Type,
	}
}

func CreateToModel(create *Create) Exercise {
	return Exercise{
		Title:       create.Title,
		Description: create.Description,
		Video:       create.Video,
		Difficulty:  create.Difficulty,
		Member:      create.Member,
		Type:        create.Type,
	}
}
func PublicToModel(public *Public, exerciseId uint) *Exercise {
	return &Exercise{
		ID:          exerciseId,
		Title:       public.Title,
		Description: public.Description,
		Video:       public.Video,
		Difficulty:  public.Difficulty,
		Member:      public.Member,
		Type:        public.Type,
	}
}
