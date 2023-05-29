package objectiveModel

import (
	"time"

	timeUtils "github.com/kilianp07/MuscleApp/utils/time"
	"gorm.io/gorm"
)

type Objective struct {
	gorm.Model
	ID          uint       `json:"id" gorm:"primaryKey"`
	UserID      uint       `json:"user_id"`
	Title       string     `json:"title"`
	Weight      float32    `json:"weight"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date"`
}

type Public struct {
	Title       string  `json:"title"`
	Weight      float32 `json:"weight"`
	Description string  `json:"description"`
	Date        int64   `json:"date"`
}

func ModelToPublic(objective *Objective) *Public {
	return &Public{
		Title:       objective.Title,
		Weight:      objective.Weight,
		Description: objective.Description,
		Date:        timeUtils.TimeToTimestamp(objective.Date),
	}
}

func PublicToModel(public *Public, userid uint) *Objective {
	return &Objective{
		UserID:      userid,
		Title:       public.Title,
		Weight:      public.Weight,
		Description: public.Description,
		Date:        timeUtils.TimestampToTime(public.Date),
	}

}
