package weightModel

import (
	"time"

	timeUtils "github.com/kilianp07/MuscleApp/utils/time"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID     uint       `json:"id" gorm:"primary_key"`
	UserID uint       `json:"user_id"`
	Date   *time.Time `json:"date" binding:"required"`
	Value  float64    `json:"value" binding:"required"`
}

type Create struct {
	Date  int64   `json:"date" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

type Public struct {
	Date  int64   `json:"date" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

func ModelToPublic(weight *Model) *Public {
	return &Public{
		Date:  timeUtils.TimeToTimestamp(weight.Date),
		Value: weight.Value,
	}
}

func PublicToModel(weight *Public, timestamp int64, userId uint) *Model {

	return &Model{
		Date:   timeUtils.TimestampToTime(int64(timestamp)),
		Value:  weight.Value,
		UserID: userId,
	}
}
