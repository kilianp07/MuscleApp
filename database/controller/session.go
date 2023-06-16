package controller

import sessionModel "github.com/kilianp07/MuscleApp/models/session"

func (c *Controller) CreateSession(session *sessionModel.Session) error {
	return c.db.Create(&session).Error
}

func (c *Controller) UpdateSession(session *sessionModel.Session) error {
	return c.db.Save(&session).Error
}

func (c *Controller) DeleteSession(id uint) error {
	return c.db.Delete(&sessionModel.Session{}, id).Error
}

func (c *Controller) GetSessions(userId uint) ([]*sessionModel.Public, error) {
	var (
		data     []*sessionModel.Session
		sessions []*sessionModel.Public
	)
	err := c.db.Where("user_id = ?", userId).Find(&data).Error

	for _, session := range data {
		sessions = append(sessions, sessionModel.ModelToPublic(session))
	}

	if err != nil {
		return sessions, err
	}
	return sessions, nil
}

func (c *Controller) GetSessionByID(id uint) (*sessionModel.Session, error) {
	var (
		session sessionModel.Session
	)
	err := c.db.First(&session, id).Error

	if err != nil {
		return &session, err
	}
	return &session, nil
}
