package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kilianp07/MuscleApp/database"
	authHandler "github.com/kilianp07/MuscleApp/handlers/auth"
	userHandler "github.com/kilianp07/MuscleApp/handlers/user"
	weightHandler "github.com/kilianp07/MuscleApp/handlers/weight"
	tokenutil "github.com/kilianp07/MuscleApp/utils/tokens"
	"gorm.io/gorm"
)

type Api struct {
	userH   *userHandler.UserHandler
	authH   *authHandler.AuthHandler
	weightH *weightHandler.WeightHandler
	db      *gorm.DB
}

func NewApi() *Api {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	return &Api{
		userH:   userHandler.NewUserHandler(db),
		authH:   authHandler.NewAuthHandler(db),
		weightH: weightHandler.NewWeightHandler(db),
		db:      db,
	}
}

func (api *Api) StartApi() {
	r := gin.Default()
	r.Use(CORS())
	api.createGroups(r)
	err := r.Run(":" + os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}
}

func (api *Api) createGroups(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/:id", api.userH.GetUserByID, tokenutil.JwtAuthMiddleware())
		user.POST("/", api.userH.CreateUser)
		user.PUT("/:id", api.userH.UpdateUser, tokenutil.JwtAuthMiddleware())
		user.DELETE("/:id", api.userH.DeleteUser)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/login", api.authH.Login)
		auth.POST("/refresh", tokenutil.JwtAuthMiddleware(), api.authH.RefreshToken)
	}

	weight := r.Group("/weight", tokenutil.JwtAuthMiddleware())
	{
		weight.POST("/", api.weightH.CreateWeight)
		weight.GET("/latest", api.weightH.GetLatestWeight)
		weight.GET("/", api.weightH.GetWeights)
		weight.GET("/:start/:end", api.weightH.GetWeightsBetween)
		weight.DELETE("/:date", api.weightH.DeleteWeight)

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
