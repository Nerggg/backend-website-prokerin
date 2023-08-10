package api

import (
	"backend-prokerin/controller"
	"backend-prokerin/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupApiRoute(router *gin.Engine) {

	// router.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization, content-type"}
	router.Use(cors.New(config))
	router.Use(gin.Recovery())

	setupPublicRoute(router)
	setupProtectedRoute(router)

}

func setupPublicRoute(router *gin.Engine) {
	publicRoutes := router.Group("/api/auth")
	// publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.POST("/register", controller.RegisterUserAccount)
	// publicRoutes.POST("/init", controller.InitUser) //untuk buat user pertama kali
}
func setupProtectedRoute(router *gin.Engine) {
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/proker", controller.AddProker)
	protectedRoutes.GET("/proker/:id", controller.GetProker)
	protectedRoutes.GET("/proker", controller.GetAllProker)
	protectedRoutes.GET("/proker/:id/like", controller.LikeProker)
	protectedRoutes.GET("/proker/:id/unlike", controller.UnLikeProker)
	protectedRoutes.DELETE("/proker/:id", controller.DeleteProker)

	protectedRoutes.POST("/comment/:proker", controller.AddComment)
	protectedRoutes.GET("/comment/:proker", controller.GetComment)
	protectedRoutes.GET("/comment", controller.GetAllComment)
	protectedRoutes.GET("/comment/:proker/:id/like", controller.LikeComment)
	protectedRoutes.GET("/comment/:proker/:id/unlike", controller.UnLikeComment)
	protectedRoutes.DELETE("/comment/:id", controller.DeleteComment)

}
