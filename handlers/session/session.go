package sessionHandler

import (
	"net/http"
	"strconv"

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
		data    sessionModel.Create
		session sessionModel.Session
		userId  uint
		err     error
	)

	userId, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session = sessionModel.CreateToModel(data, userId)

	if err := handler.controller.CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessionModel.ModelToPublic(&session))
}

func (handler *SessionHandler) GetSessionByID(c *gin.Context) {
	var (
		id     string
		userID uint
		err    error
	)

	id = c.Param("id")

	userID, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := handler.controller.GetSessionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if session.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this session"})
		return
	}

	c.JSON(http.StatusOK, sessionModel.ModelToPublic(session))
}

func (handler *SessionHandler) UpdateSession(c *gin.Context) {
	var (
		session sessionModel.Public
		updated *sessionModel.Session
		err     error
		idS     string
		id      int
	)

	if err = c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idS = c.Param("id")
	id, err = strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionM := sessionModel.PublicToModel(session)
	sessionM.ID = uint(id)

	if err := handler.controller.UpdateSession(&sessionM); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updated, err = handler.controller.GetSessionByID(idS); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sessionModel.ModelToPublic(updated))
}

func (handler *SessionHandler) DeleteSession(c *gin.Context) {
	id := c.Param("id")

	userId, err := gincontext.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := handler.controller.GetSessionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if session.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this session"})
		return
	}

	if err := handler.controller.DeleteSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}

func (handler *SessionHandler) GetAllSessions(c *gin.Context) {

	var (
		userID uint
		err    error
	)

	userID, err = gincontext.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := handler.controller.GetAllSessions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	public := make([]sessionModel.Public, len(session))
	for i, s := range session {
		public[i] = sessionModel.ModelToPublic(&s)
	}

	c.JSON(http.StatusOK, public)
}
