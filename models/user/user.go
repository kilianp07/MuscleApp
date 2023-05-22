package userModel

type Model struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"  binding:"required"`
	Email    string `json:"email" gorm:"unique"  binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
	Username string `json:"username" gorm:"unique"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Public struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ModelToPublic(user *Model) *Public {
	return &Public{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Username,
		Email:    user.Email,
	}
}
