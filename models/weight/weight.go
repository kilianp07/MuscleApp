package weightModel

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID     uint       `json:"id" gorm:"primary_key"`
	UserID uint       `json:"user_id"`
	Date   *time.Time `json:"date"`
	Value  float64    `json:"value"`
}

type Public struct {
	Date  *time.Time `json:"date"`
	Value float64    `json:"value"`
}
