package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/handlers"
	"github.com/omgupta1608/chatex/server/pkg/middleware/jwt"
)

func InitBaseRoute(engine *gin.Engine) {
	engine.GET("/", handlers.BaseRouteHandler)
}

func InitPublicRoutes(router *gin.RouterGroup) {
	router.POST("/register", handlers.RegisterRouteHandler)
	router.POST("/register/verify", handlers.UserVerificationRouteHandler)
	router.POST("/login", handlers.LoginRouteHandler)
}

func InitPrivateRoutes(router *gin.RouterGroup) {
	router.Use(jwt.JWTAuthMiddleware())
	// Get User By Id Route
	router.GET("/user/:uid", handlers.GetUserById)
}
