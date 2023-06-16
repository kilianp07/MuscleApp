package exerciseModel

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

type Public struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
}

func ModelToPublic(exercise *Exercise) *Public {
	return &Public{
		ID:          exercise.ID,
		Title:       exercise.Title,
		Description: exercise.Description,
		Image:       exercise.Image,
		Video:       exercise.Video,
	}
}

func PublicToModel(public *Public, exerciseId uint) *Exercise {
	return &Exercise{
		ID:          exerciseId,
		Title:       public.Title,
		Description: public.Description,
		Image:       public.Image,
		Video:       public.Video,
	}
}
