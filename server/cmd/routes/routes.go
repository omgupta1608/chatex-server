package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/cmd/config"
	"github.com/omgupta1608/chatex/server/cmd/handlers"
	"github.com/omgupta1608/chatex/server/pkg/middleware/jwt"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

func InitBaseRoute(engine *gin.Engine) {
	engine.GET("/", handlers.BaseRouteHandler)
}

var jwtService = jwt.New(&jwt.JWTService{
	SigningMethodStr: "HS256",
	SecretKey:        config.GetJwtSecret(),
	CookieName:       "jwt",
})

func InitPublicRoutes(router *gin.RouterGroup) {
	/**
	* ! Only for dev testing. Remove for production
	 */
	router.GET("/gettoken", func(c *gin.Context) {
		var user types.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		fmt.Println("user ", user)
		tokenString, err := jwtService.GenerateTokenString(user.Uid, user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		js := jwtService
		c.SetCookie(
			js.CookieName,
			tokenString,
			js.CookieMaxAge,
			js.CookiePath,
			js.CookieDomain,
			js.SecureCookie,
			js.HttpOnlyCookie,
		)

		c.JSON(http.StatusOK, tokenString)
	})
}

func InitPrivateRoutes(router *gin.RouterGroup) {
	router.Use(jwt.JWTAuthMiddleware(jwtService))

	/**
	* ! Only for dev testing. Remove for production
	 */
	router.GET("/", func(c *gin.Context) {
		user, ok := c.MustGet("User").(types.User)
		json, err := json.Marshal(user)

		if !ok || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "no user",
			})
		} else {
			c.JSON(http.StatusOK, string(json))
		}
	})
}
