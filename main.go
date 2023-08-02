package main

import (
	"backend-prokerin/api"
	"backend-prokerin/initialization"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/iand/logfmtr"
)

func main() {

	// mode := os.Getenv("MODE")
	// if mode == "PROD" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	api.SetupApiRoute(router)
	var port string = os.Getenv("PORT")

	initialization.CreateSuperAdmin()

	if port == "" {
		port = "8001"
	}
	router.Run(":" + port)
	// logger
	var logger = logfmtr.New().WithName("main")
	logger.Info("Server running on port " + port)

}
