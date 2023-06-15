package weightHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	gincontext "github.com/kilianp07/MuscleApp/utils/gin_context"
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
		data   weightModel.Public
		weight weightModel.Weight
		err    error
	)

	if weight.UserID, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Paste the data into the weight object
	weight.Date = timeUtils.TimestampToTime(data.Date)
	weight.Value = data.Value

	// Save the weight to the database
	if err = handler.controller.CreateWeight(&weight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, weight)
}

func (handler *WeightHandler) DeleteWeight(c *gin.Context) {
	var (
		err    error
		userId uint
	)

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	weightDate := c.Param("date")
	weightDateInt, err := strconv.Atoi(weightDate)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot read weight date"})
		return
	}

	// Delete the weight from the database
	if err = handler.controller.DeleteWeightByDate(userId, weightDateInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Weight deleted"})

}

func (handler *WeightHandler) UpdateWeight(c *gin.Context) {
	var (
		weight weightModel.Public
		err    error
		userId uint
	)

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&weight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the weight in the database
	if err = handler.controller.UpdateWeightByDate(userId, &weight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, weight)
}

func (handler *WeightHandler) GetLatestWeight(c *gin.Context) {
	var (
		weightResult *weightModel.Weight
		err          error
		userId       uint
	)
	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the weight with the given ID from the database
	if weightResult, err = handler.controller.GetLatestWeight(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weightResult)
}

func (handler *WeightHandler) GetWeights(c *gin.Context) {
	var (
		weights []*weightModel.Public
		err     error
		userId  uint
	)

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the weights with the given ID from the database
	if weights, err = handler.controller.GetWeights(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weights)
}

func (handler *WeightHandler) GetInitialWeight(c *gin.Context) {
	var (
		weightResult *weightModel.Weight
		err          error
		userId       uint
	)
	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the weight with the given ID from the database
	if weightResult, err = handler.controller.GetInitialWeight(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weightResult)
}

// Get between two timestamps (start and end)
func (handler *WeightHandler) GetWeightsBetween(c *gin.Context) {
	var (
		weights []*weightModel.Public
		err     error
		userId  uint
	)

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	startDate, err := strconv.Atoi(c.Param("start"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot recognize start date"})
		return
	}
	endDate, err := strconv.Atoi(c.Param("end"))
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot recognize end date"})
		return
	}

	// Retrieve the weights with the given ID from the database
	if weights, err = handler.controller.GetWeightsBetween(userId, startDate, endDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, weights)
}
