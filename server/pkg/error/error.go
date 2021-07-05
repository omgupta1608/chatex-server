package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getError(e error) string {
	if e != nil {
		return e.Error()
	}
	return ""
}

func SendError(c *gin.Context, code int, err error) {
	switch {
	case code >= 400 && code < 500:
		c.JSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "Bad Request",
			"error":   getError(err),
		})
	case code >= 500 && code < 600:
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": "Something went wrong!",
			"error":   getError(err),
		})
	}
}
