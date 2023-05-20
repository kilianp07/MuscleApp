package weightHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	"gorm.io/gorm"
)

type WeightHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewWeightHandler(db *gorm.DB) *WeightHandler {
	return &WeightHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *WeightHandler) CreateWeight(c *gin.Context) {
	var (
		weight weightModel.Model
		err    error
	)

	// Get the user from the refresh token
	userId, exists := c.Get("x-user-id")
	if !exists {
		c.JSON(500, gin.H{"error": "User id missing"})
		return
	}
	// Paste it into the weight object
	userIdstring := userId.(string)
	userIdUint, err := strconv.ParseUint(userIdstring, 10, 32)
	if err != nil {
		c.JSON(500, gin.H{"error": "User id missing"})
		return
	}
	weight.UserID = uint(userIdUint)

	if err = c.ShouldBindJSON(&weight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the weight to the database
	if err = handler.controller.CreateWeight(&weight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, weight)
}

func (handler *WeightHandler) GetLatestWeight(c *gin.Context) {
	var (
		weightResult *weightModel.Model
		err          error
	)

	// Get the user from the refresh token
	userId, exists := c.Get("x-user-id")
	if !exists {
		c.JSON(500, gin.H{"error": "User id missing"})
		return
	}
	userIdstring := userId.(string)
	userIdUint, err := strconv.ParseUint(userIdstring, 10, 32)
	if err != nil {
		c.JSON(500, gin.H{"error": "User id missing"})
		return
	}

	// Retrieve the weight with the given ID from the database
	if weightResult, err = handler.controller.GetLatestWeight(uint(userIdUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weightResult)
}
