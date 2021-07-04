package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

// sends an unauthorized http response
func sendUnauthorized(c *gin.Context, message string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"error":   err,
	})
}

// middleware to authenticate using jwt token
func JWTAuthMiddleware(service *JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get jwt stored in cookie
		tokenString, err := c.Cookie(service.CookieName)
		if err != nil {
			sendUnauthorized(c, "Please Login", err)
		}

		// parse jwt token
		authCustomClaims, err := service.ParseToken(tokenString)
		if err != nil {
			sendUnauthorized(c, "Please Login again. Bad Token", err)
			c.Abort()
			return
		}

		// store user data in this context
		user := types.User{Uid: authCustomClaims.Uid, Email: authCustomClaims.Email}
		c.Set("User", user)
		c.Next()
	}
}
