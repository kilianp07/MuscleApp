package weightHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	timeUtils "github.com/kilianp07/MuscleApp/utils/time"
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
		data   weightModel.Create
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
		c.JSON(500, gin.H{"error": "Cannot read user id"})
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Paste the data into the weight object
	weight.UserID = uint(userIdUint)
	weight.Date = timeUtils.TimestampToTime(data.Date)
	weight.Value = data.Value

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

func (handler *WeightHandler) GetWeights(c *gin.Context) {
	var (
		weights []*weightModel.Public
		err     error
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

	// Retrieve the weights with the given ID from the database
	if weights, err = handler.controller.GetWeights(uint(userIdUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weights)
}

// Get between two timestamps (start and end)
func (handler *WeightHandler) GetWeightsBetween(c *gin.Context) {
	var (
		weights []*weightModel.Public
		err     error
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

	startDate, err := strconv.ParseInt(c.Param("start"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot recognize start date"})
		return
	}
	endDate, err := strconv.ParseInt(c.Param("end"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot recognize end date"})
		return
	}

	// Retrieve the weights with the given ID from the database
	if weights, err = handler.controller.GetWeightsBetween(uint(userIdUint), int(startDate), int(endDate)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weights)
}
