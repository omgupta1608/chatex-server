package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/config"
	"github.com/omgupta1608/chatex/server/cmd/routes"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)  // Uncomment when in prod

	router := gin.New()

	// Add CORS middleware

	publicRoutes := router.Group("api/" + config.GetApiVersion() + "/")
	privateRoutes := router.Group("api/" + config.GetApiVersion() + "/")

	// attach auth middleware to private routes

	routes.InitBaseRoute(router)
	routes.InitPrivateRoutes(privateRoutes)
	routes.InitPublicRoutes(publicRoutes)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Print("Server started at PORT => " + PORT + "\n")
	router.Run(":" + PORT)
}
