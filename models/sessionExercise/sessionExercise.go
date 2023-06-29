package sessionexerciseModel

import "gorm.io/gorm"

type SessionExercise struct {
	gorm.Model
	ID         uint    `json:"id" gorm:"primary_key"`
	SessionID  uint    `json:"session_id" gorm:"type:bigint unsigned"`
	Name       string  `json:"name"`
	Serie      uint    `json:"serie"`
	Repetition uint    `json:"repetition"`
	RestTime   float32 `json:"rest_time"`
	Method     string  `json:"method"`
}
type Public struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Serie      uint    `json:"serie"`
	Repetition uint    `json:"repetition"`
	RestTime   float32 `json:"rest_time"`
	Method     string  `json:"method"`
}
type Create struct {
	SessionID  uint    `json:"session_id" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Serie      uint    `json:"serie" binding:"required"`
	Repetition uint    `json:"repetition" binding:"required"`
	RestTime   float32 `json:"rest_time" binding:"required"`
	Method     string  `json:"method" binding:"required"`
}

func ModelToPublic(session SessionExercise) Public {
	return Public{
		ID:         session.ID,
		Name:       session.Name,
		Serie:      session.Serie,
		Repetition: session.Repetition,
		RestTime:   session.RestTime,
		Method:     session.Method,
	}
}
func CreateToModel(session Create) SessionExercise {
	return SessionExercise{
		SessionID:  session.SessionID,
		Name:       session.Name,
		Serie:      session.Serie,
		Repetition: session.Repetition,
		RestTime:   session.RestTime,
		Method:     session.Method,
	}
}

func PublicToModel(session Public) SessionExercise {
	return SessionExercise{
		ID:         session.ID,
		Name:       session.Name,
		Serie:      session.Serie,
		Repetition: session.Repetition,
		RestTime:   session.RestTime,
		Method:     session.Method,
	}
}
