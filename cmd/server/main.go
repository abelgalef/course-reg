package main

import (
	"flag"
	"net/http"

	"github.com/abelgalef/course-reg/pkg/handlers/rest"
	"github.com/abelgalef/course-reg/pkg/middlewares"
	"github.com/abelgalef/course-reg/pkg/services"
	"github.com/abelgalef/course-reg/pkg/storage"
	"github.com/abelgalef/course-reg/pkg/storage/repos"
	"github.com/gin-gonic/gin"
)

func setupAuthRoute(r *gin.Engine, handler rest.AuthHandler) *gin.Engine {
	auth := r.Group("/auth")
	{
		auth.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "wass goood?")
		})

		auth.POST("/login", func(ctx *gin.Context) {
			res, err := handler.Login(ctx)
			if err != nil {
				ctx.JSON(err[0].Code, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		auth.POST("/register", func(ctx *gin.Context) {
			res, err := handler.Register(ctx)
			if err != nil {
				ctx.JSON(err[0].Code, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})
	}

	return r
}

func main() {

	f := flag.Bool("sqlite", false, "Set this flag to use a sqlite database; The default database is MySQL.")

	// SETUP STORAGE
	db := storage.NewDatabaseConnection(*f)

	// SETUP REPO
	userRepo := repos.NewUserRepo(db)

	// SETUP SERVICES
	JWTservice := services.NewJWTService()
	authService := services.NewAuthService(userRepo, JWTservice)

	// REST HANDLERS
	handler := rest.NewAuthHandler(authService)

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r = setupAuthRoute(r, handler)
	// r = setupDeptRouter(r, handler)
	r.Run(":8080")
}
