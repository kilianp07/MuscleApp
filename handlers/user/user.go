package userHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	userModel "github.com/kilianp07/MuscleApp/models/user"
	"gorm.io/gorm"
)

type UserHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var (
		user userModel.User
		err  error
	)

	// Make sure there is no other user with the same email or username

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the user to the database
	if err = handler.controller.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (handler *UserHandler) GetUserByID(c *gin.Context) {
	var (
		userResult *userModel.User
		err        error
	)
	id := c.Param("id")

	// Retrieve the user with the given ID from the database
	if userResult, err = handler.controller.GetUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, userModel.ModelToPublic(userResult))
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var (
		user *userModel.User
		err  error
	)

	// Check if the user exists
	if user, err = handler.controller.GetUserByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Retrieve the user with the given ID from the database
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user in the database
	if err := handler.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, userModel.ModelToPublic(user))
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user userModel.User

	// Retrieve the user with the given ID from the database
	if err := handler.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}
	// Delete the user from the database

	if err := handler.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
