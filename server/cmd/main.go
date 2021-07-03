package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)  // Uncomment when in prod

	r := gin.New()

	r.GET("/a", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Chatex!",
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Print("Server started at PORT => " + PORT + "\n")
	r.Run(":" + PORT)
}
