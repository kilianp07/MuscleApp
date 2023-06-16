package sessionHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database/controller"
	sessionModel "github.com/kilianp07/MuscleApp/models/session"
	gincontext "github.com/kilianp07/MuscleApp/utils/gin_context"
	"gorm.io/gorm"
)

type SessionHandler struct {
	db         *gorm.DB
	controller *controller.Controller
}

func NewSessionHandler(db *gorm.DB) *SessionHandler {
	return &SessionHandler{
		db:         db,
		controller: controller.NewController(db),
	}
}

func (handler *SessionHandler) CreateSession(c *gin.Context) {

	var (
		err     error
		userId  uint
		session sessionModel.Public
	)

	if userId, err = gincontext.GetUserId(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.controller.CreateSession(sessionModel.PublicToModel(&session, userId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func (handler *SessionHandler) GetSessionByID(c *gin.Context) {
	var (
		sessionResult *sessionModel.Session
		err           error
		id            uint
	)

	id = c.Param("id")

	if sessionResult, err = handler.controller.GetSessionByID(sessionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, sessionModel.ModelToPublic(sessionResult))
}
