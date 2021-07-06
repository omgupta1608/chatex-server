package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/firebase"
)

// Handler function for Get User By Id route
// @param {string} uid - Id of the User
// /api/v/user/:uid
func GetUserById(c *gin.Context) {
	uid, ok := c.Params.Get("uid")

	if !ok {
		exception.SendError(c, 403, errors.New("no user id provided"))
		return
	}

	dbsnapshot, err := firebase.Client.Collection("Users").Doc(uid).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, 500, nil)
		return
	}

	user := dbsnapshot.Data()

	delete(user, "Password")

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "",
		"error":   nil,
	})
}
