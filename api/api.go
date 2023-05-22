package api

import (
	"os"

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
	r.Use(CORS())
	api.createGroups(r)
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

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
