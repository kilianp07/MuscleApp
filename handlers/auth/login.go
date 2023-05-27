package authHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	"github.com/kilianp07/MuscleApp/utils/auth"
	gincontext "github.com/kilianp07/MuscleApp/utils/gin_context"
	tokenutil "github.com/kilianp07/MuscleApp/utils/tokens"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *AuthHandler) Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.JSON(400, gin.H{"error": "Email or password missing"})
		return
	}

	// Check if the user exists
	user, err := handler.controller.GetUserByEmail(email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Check if the password is correct
	if !auth.CheckPasswordHash(password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a token
	token, err := tokenutil.CreateAccessToken(user, 1)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Generate a refresh token
	refreshToken, err := tokenutil.CreateRefreshToken(user, 24)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token, "refreshToken": refreshToken})

}

func (handler *AuthHandler) RefreshToken(c *gin.Context) {
	var (
		userId uint
		err    error
	)

	refreshToken := c.Request.Header.Get("Authorization")
	if refreshToken == "" {
		c.JSON(400, gin.H{"error": "Refresh token missing"})
		return
	}

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	user, err := handler.controller.GetUserByID(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Generate a new access token
	token, err := tokenutil.CreateAccessToken(user, 1)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})

}
