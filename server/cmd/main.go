package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/config"
	"github.com/omgupta1608/chatex/server/cmd/routes"
	database "github.com/omgupta1608/chatex/server/pkg/firebase"
)

func main() {
	buildFlags := config.GetBuildFlags()

	if buildFlags.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Initialize Firestore Database
	database.InitFirestore()

	// Add CORS middleware

	publicRoutes := router.Group("api/" + config.GetApiVersion() + "/")
	privateRoutes := router.Group("api/" + config.GetApiVersion() + "/")

	// attach auth middleware to private routes

	routes.InitBaseRoute(router)
	routes.InitPrivateRoutes(privateRoutes)
	routes.InitPublicRoutes(publicRoutes)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = buildFlags.Port
	}

	fmt.Print("Server started at PORT => " + PORT + "\n")
	router.Run(":" + PORT)
}
