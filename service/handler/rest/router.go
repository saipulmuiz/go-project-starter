package rest

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/go-project-starter/service"
	middleware "github.com/saipulmuiz/go-project-starter/service/middleware"
)

type Handler struct {
	userUsecase     service.UserUsecase
	categoryUsecase service.CategoryUsecase
}

func CreateHandler(
	userUsecase service.UserUsecase,
	categoryUsecase service.CategoryUsecase,
) *gin.Engine {
	obj := Handler{
		userUsecase:     userUsecase,
		categoryUsecase: categoryUsecase,
	}

	r := gin.Default()

	gin.SetMode(gin.DebugMode)
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(middleware.LoggingMiddleware())
	r.Use(gin.Recovery())
	r.Use(timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "success"})
		}),
	))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST, PUT, PATCH, DELETE, OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	publicRouter := r.Group("/v1")
	publicRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "welcome, to golang project starter",
		})
	})
	publicRouter.POST("/register", obj.Register)
	publicRouter.POST("/login", obj.login)

	authRouter := publicRouter.Group("/")
	authRouter.Use(middleware.AuthMiddleware())
	{
		// Categories
		authRouter.GET("/categories", obj.GetCategories)
		authRouter.POST("/categories", obj.CreateCategory)
		authRouter.PUT("/categories/:categoryId", obj.UpdateCategory)
		authRouter.DELETE("/categories/:categoryId", obj.DeleteCategory)
	}

	return r
}
