package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// sends an unauthorized http response
func sendUnauthorized(c *gin.Context, message string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
		"error":   err.Error(),
	})
	c.Abort()
}

// middleware to authenticate using jwt token
func JWTAuthMiddleware() gin.HandlerFunc {
	const BEARER_SCHEMA = "Bearer"

	return func(c *gin.Context) {
		// get jwt stored in cookie
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			sendUnauthorized(c, "Please login again. Bad Token", nil)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		// parse jwt token
		user, err := ParseToken(tokenString)
		if err != nil {
			sendUnauthorized(c, "Please Login again. Bad Token", err)
			return
		}

		// store user data in this context
		c.Set("User", user)
		c.Next()
	}
}
