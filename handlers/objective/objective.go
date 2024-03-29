package objectiveHandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	objectiveModel "github.com/kilianp07/MuscleApp/models/objective"
	gincontext "github.com/kilianp07/MuscleApp/utils/gin_context"
	"gorm.io/gorm"
)

type ObjectiveHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewObjectiveHandler(db *gorm.DB) *ObjectiveHandler {
	return &ObjectiveHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *ObjectiveHandler) CreateObjective(c *gin.Context) {

	var (
		data      objectiveModel.Public
		objective objectiveModel.Objective
		err       error
	)

	if objective.UserID, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objective = *objectiveModel.PublicToModel(&data, objective.UserID)

	if err = handler.controller.CreateObjective(&objective); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objectiveModel.ModelToPublic(&objective))
}

func (handler *ObjectiveHandler) GetObjectives(c *gin.Context) {

	var (
		objectives []objectiveModel.Public
		userId     uint
		err        error
	)

	userId, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if objectives, err = handler.controller.GetAllObjectives(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objectives)
}

func (handler *ObjectiveHandler) UpdateObjective(c *gin.Context) {

	var (
		data      objectiveModel.Public
		objective *objectiveModel.Objective
		err       error
		id        int
	)

	objective = &objectiveModel.Objective{}

	idS := c.Param("id")
	if id, err = strconv.Atoi(idS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if objective.UserID, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objective = objectiveModel.PublicToModel(&data, objective.UserID)

	if err = handler.controller.UpdateObjective(objective, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if objective, err = handler.controller.GetObjectiveByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, objectiveModel.ModelToPublic(objective))
}

func (handler *ObjectiveHandler) DeleteObjective(c *gin.Context) {

	var (
		objective *objectiveModel.Objective
		err       error
		id        int
	)

	idS := c.Param("id")
	if id, err = strconv.Atoi(idS); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if objective, err = handler.controller.GetObjectiveByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = handler.controller.DeleteObjective(objective); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Objective deleted successfully"})
}
