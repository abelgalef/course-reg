package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/abelgalef/course-reg/pkg/handlers/rest"
	"github.com/abelgalef/course-reg/pkg/middlewares"
	"github.com/abelgalef/course-reg/pkg/models"
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

func setupRoleRoute(r *gin.Engine, handler rest.RoleHandler, authMiddleware gin.HandlerFunc, corsMiddleware gin.HandlerFunc) *gin.Engine {
	role := r.Group("/role", authMiddleware, corsMiddleware)
	{
		role.GET("/", func(ctx *gin.Context) {
			res, err := handler.GetRoles(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		role.GET("/:id", func(ctx *gin.Context) {
			res, err := handler.GetRole(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		role.POST("/", func(ctx *gin.Context) {
			res, err := handler.AddNewRole(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		role.POST("/give-perm/:role_id/:right_id", func(ctx *gin.Context) {
			err := handler.GivePerm(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": nil})
			}
		})

		role.POST("/revok-perm/:role_id/:right_id", func(ctx *gin.Context) {
			err := handler.RevokPerm(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": nil})
			}
		})

	}

	return r
}

func setupRightsRoute(r *gin.Engine, handler rest.RightHanler, authMiddleware gin.HandlerFunc, corsMiddleware gin.HandlerFunc) *gin.Engine {
	right := r.Group("/right", authMiddleware, corsMiddleware)
	{
		right.GET("/", func(ctx *gin.Context) {
			res, err := handler.GetRights()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})
	}

	return r
}

func setupDeptRouter(r *gin.Engine, handler rest.DeptHandler, authMiddleware gin.HandlerFunc) *gin.Engine {
	dept := r.Group("/dept", authMiddleware)
	{
		dept.GET("/", func(ctx *gin.Context) {
			res, err := handler.GetAllDepts(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		dept.GET("/:id", func(ctx *gin.Context) {
			res, err := handler.GetDept(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		dept.POST("/", func(ctx *gin.Context) {
			res, err := handler.AddDept(ctx)
			if err != nil {
				ctx.JSON(err[0].Code, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		dept.PUT("/:id", func(ctx *gin.Context) {
			res, err := handler.UpdateDept(ctx)
			if err != nil {
				ctx.JSON(err[0].Code, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		dept.DELETE("/:id", func(ctx *gin.Context) {
			err := handler.DeleteDept(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, gin.H{})
			}
		})

	}

	return r
}

func setupCourserouter(r *gin.Engine, handler rest.CourseHandler, authMiddleware gin.HandlerFunc) *gin.Engine {
	course := r.Group("/course", authMiddleware)
	{
		course.GET("/", func(ctx *gin.Context) {
			res, err := handler.GetAllCourses(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		course.GET("/:id", func(ctx *gin.Context) {
			res, err := handler.GetDeptCourses(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		course.POST("/", func(ctx *gin.Context) {
			res, err := handler.CreateCourse(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		course.PATCH("/:id", func(ctx *gin.Context) {
			res, err := handler.UpdateCourse(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		course.POST("/take-course/:id", func(ctx *gin.Context) {
			err := handler.TakeCourse(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, nil)
			}
		})

		course.GET("/all-history", func(ctx *gin.Context) {
			res, err := handler.GetAllMyCourseHistory(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		course.GET("/current-history", func(ctx *gin.Context) {
			res, err := handler.GetMyCurrentCourseHistory(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		return r
	}
}

func setupConstRoutes(r *gin.Engine, handler rest.ConstHandler, authMiddleware gin.HandlerFunc) *gin.Engine {
	con := r.Group("/const", authMiddleware)
	{
		con.GET("/", func(ctx *gin.Context) {
			res, err := handler.GetAllConstraintHistory(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		con.GET("/:id", func(ctx *gin.Context) {
			res, err := handler.GetDeptConstraintHistory(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		con.GET("/current/:id", func(ctx *gin.Context) {
			res, err := handler.GetCurrentConstraint(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		con.POST("/", func(ctx *gin.Context) {
			res, err := handler.CreateConstraint(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
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
	roleRepo := repos.NewRoleRepository(db)
	rightRepo := repos.NewRightRepository(db)
	deptRepo := repos.NewDeptRepo(db)
	constRepo := repos.NewConstRepository(db)
	courseRepo := repos.NewCourseRepo(db)

	// SETUP SERVICES
	JWTservice := services.NewJWTService()
	authService := services.NewAuthService(userRepo, JWTservice)
	roleService := services.NewRoleService(rightRepo, roleRepo)
	deptService := services.NewDeptService(deptRepo, rightRepo)
	constService := services.NewConstService(constRepo, rightRepo)
	courseService := services.NewCourseService(courseRepo, rightRepo, constRepo)
	rightService := services.NewRightService(rightRepo)

	// REST HANDLERS
	authHandler := rest.NewAuthHandler(authService)
	roleHandler := rest.NewRoleHandler(roleService)
	deptHandler := rest.NewDeptHandler(deptService)
	constHanlder := rest.NewConstHandler(constService)
	courseHandler := rest.NewCourseHandler(courseService)
	rightHandler := rest.NewRightHandler(rightService)

	authMiddleware := middlewares.CheckAuth(JWTservice, userRepo)
	corsMiddleware := middlewares.CORSMiddleware()

	r := gin.Default()
	r.Use(corsMiddleware)

	r.GET("/user", authMiddleware, func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, &models.LoginResponse{User: *(ctx.MustGet("user").(*models.User)), Token: strings.Split(ctx.GetHeader("Authorization"), " ")[1]})
	})

	r = setupAuthRoute(r, authHandler)
	r = setupRoleRoute(r, roleHandler, authMiddleware, corsMiddleware)
	r = setupDeptRouter(r, deptHandler, authMiddleware)
	r = setupConstRoutes(r, constHanlder, authMiddleware)
	r = setupCourserouter(r, courseHandler, authMiddleware)
	r = setupRightsRoute(r, rightHandler, authMiddleware, corsMiddleware)

	r.Run(":8080")
}
