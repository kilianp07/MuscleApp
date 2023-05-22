package api

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database"
	authHandler "github.com/kilianp07/MuscleApp/handlers/auth"
	userHandler "github.com/kilianp07/MuscleApp/handlers/user"
	tokenutil "github.com/kilianp07/MuscleApp/utils/tokens"
	"gorm.io/gorm"
)

type Api struct {
	userH *userHandler.UserHandler
	authH *authHandler.AuthHandler
	db    *gorm.DB
}

func NewApi() *Api {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	return &Api{
		userH: userHandler.NewUserHandler(db),
		authH: authHandler.NewAuthHandler(db),
		db:    db,
	}
}

func (api *Api) StartApi() {
	r := gin.Default()
	api.createGroups(r)
	r.Use(cors.Default())
	r.Run(":" + os.Getenv("API_PORT"))
}

func (api *Api) createGroups(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/:id", api.userH.GetUserByID)
		user.POST("/", api.userH.CreateUser)
		user.PUT("/:id", api.userH.UpdateUser)
		user.DELETE("/:id", api.userH.DeleteUser)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", api.authH.Login)
		auth.POST("/refresh", tokenutil.JwtAuthMiddleware(), api.authH.RefreshToken)
	}
}
