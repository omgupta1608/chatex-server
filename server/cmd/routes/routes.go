package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/handlers"
	"github.com/omgupta1608/chatex/server/pkg/middleware/jwt"
	"github.com/omgupta1608/chatex/server/pkg/socket"
)

func InitBaseRoute(engine *gin.Engine) {
	engine.GET("/", handlers.BaseRouteHandler)
}

func InitPublicRoutes(router *gin.RouterGroup) {
	router.POST("/register", handlers.RegisterRouteHandler)
	router.POST("/register/verify", handlers.UserVerificationRouteHandler)
	router.POST("/login", handlers.LoginRouteHandler)

	// TODO: Make the socket connection private, to allow only authenticated connections
	router.GET("/ws", func(c *gin.Context) {
		socket.SocketHandler(c.Writer, c.Request)
	})
}

func InitPrivateRoutes(router *gin.RouterGroup) {
	router.Use(jwt.JWTAuthMiddleware())
	// Get User By Id Route
	router.GET("/user/:uid", handlers.GetUserById)
	// edit user profile route
	router.POST("/user/edit/:uid", handlers.EditUserProfile)
	// change user-password route
	router.POST("/user/change-password/:uid", handlers.ChangePassword)
	// delete user account route
	router.DELETE("/user/delete-account/:uid", handlers.DeleteUserById)
}
