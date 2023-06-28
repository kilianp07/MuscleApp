package exerciseHandler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	exerciseModel "github.com/kilianp07/MuscleApp/models/Exercise"
	gincontext "github.com/kilianp07/MuscleApp/utils/gin_context"
	"gorm.io/gorm"
)

type ExerciseHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewExerciseHandler(db *gorm.DB) *ExerciseHandler {
	return &ExerciseHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *ExerciseHandler) CreateExercise(c *gin.Context) {
	var (
		data     exerciseModel.Create
		exercise *exerciseModel.Exercise
		err      error
		ex       *exerciseModel.Public
	)

	_, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	exercise = exerciseModel.CreateToModel(&data)

	if err = handler.controller.CreateExercise(exercise); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if ex, err = handler.controller.GetLatestExercise(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, ex)
}

func (handler *ExerciseHandler) GetExerciseByID(c *gin.Context) {
	var (
		exercise *exerciseModel.Public
		err      error
	)

	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	exercise, err = handler.controller.GetExerciseByID(idint)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exercise)
}

func (handler *ExerciseHandler) GetSomeExercises(c *gin.Context) {
	var (
		exercises []exerciseModel.Public
		err       error
	)

	number := c.Param("number")
	numberint, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	exercises, err = handler.controller.GetSomeExercises(numberint)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exercises)
}

func (handler *ExerciseHandler) UpdateExercise(c *gin.Context) {
	var (
		data     exerciseModel.Public
		exercise *exerciseModel.Exercise
		err      error
	)

	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.controller.GetExerciseByID(idint)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	exercise = exerciseModel.PublicToModel(&data, uint(idint))

	if err = handler.controller.UpdateExercise(exercise); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exerciseModel.ModelToPublic(exercise))
}

func (handler *ExerciseHandler) DeleteExercise(c *gin.Context) {
	var (
		exercise *exerciseModel.Exercise
		err      error
	)

	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = handler.controller.GetExerciseByID(idint)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	exercise = &exerciseModel.Exercise{
		ID: uint(idint),
	}

	if err = handler.controller.DeleteExercise(exercise); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Exercise deleted"})
}
