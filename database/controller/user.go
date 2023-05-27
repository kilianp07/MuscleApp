package controller

import (
	"fmt"

	userModel "github.com/kilianp07/MuscleApp/models/user"
	"github.com/kilianp07/MuscleApp/utils/auth"
	"gorm.io/gorm"
)

func (c *Controller) GetUserByID(id uint) (*userModel.User, error) {

	var user *userModel.User
	if err := c.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user")
	}

	return user, nil
}

func (c *Controller) GetUserByEmail(email string) (*userModel.User, error) {
	var user userModel.User

	if err := c.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user")
	}

	return &user, nil
}

func (c *Controller) CreateUser(user *userModel.User) error {
	var err error

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
