package controller

import (
	sessionModel "github.com/kilianp07/MuscleApp/models/session"
	"gorm.io/gorm"
)

func (c *Controller) CreateSession(session *sessionModel.Session) error {
	return c.db.Create(session).Error
}

func (c *Controller) GetSessionByID(id string) (*sessionModel.Session, error) {
	var session sessionModel.Session
	if err := c.db.Preload("Exercises").First(&session, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (c *Controller) UpdateSession(session *sessionModel.Session) error {
	return c.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&session).Error
}

func (c *Controller) DeleteSession(session *sessionModel.Session) error {
	return c.db.Delete(&session).Error
}

func (c *Controller) GetAllSessions(userId uint) ([]sessionModel.Session, error) {
	var session []sessionModel.Session
	if err := c.db.Where("user_id = ?", userId).Preload("Exercises").Find(&session).Error; err != nil {
		return nil, err
	}
	return session, nil
}
