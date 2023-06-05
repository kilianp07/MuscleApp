package exerciseHandler

import (
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
		data     exerciseModel.Public
		exercise *exerciseModel.Exercise
		err      error
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

	exercise = &exerciseModel.Exercise{
		Title:       data.Title,
		Description: data.Description,
		Image:       data.Image,
		Video:       data.Video,
	}

	if err = handler.controller.CreateExercise(exercise); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, exerciseModel.ModelToPublic(exercise))
}
