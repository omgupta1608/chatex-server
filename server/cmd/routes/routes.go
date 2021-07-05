package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/handlers"
)

func InitBaseRoute(engine *gin.Engine) {
	engine.GET("/", handlers.BaseRouteHandler)
}

func InitPublicRoutes(router *gin.RouterGroup) {

}

func InitPrivateRoutes(router *gin.RouterGroup) {
	// Get User By Id Route
	router.GET("/user/:uid", handlers.GetUserById)
}
