package handlers

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/types"
	"github.com/rs/xid"
)

// ! not for production
var UserList []types.User

func RegisterRouteHandler(c *gin.Context) {
	// user data from request body
	var user types.User

	err := c.BindJSON(&user)
	if err != nil {
		exception.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	// TODO: sanitize client input
	// TODO: validate client input

	// initialize/generate user data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Could not hash the password"))
		return
	}
	user.Password = string(hashedPassword)
	user.Uid = xid.New().String()
	if user.About == "" {
		user.About = "Hello there!! I'm using ChatX"
	}
	// TODO: add a default profile Pic

	// TODO: store this user in redis and wait for verification
	UserList = append(UserList, user)

	// TODO: create a verification token & send through email
	// TODO: store verify token in redis

	c.JSON(http.StatusCreated, gin.H{
		"data":               user.Uid,
		"message":            "User created",
		"error":              nil,
		"verification-route": c.FullPath() + "/verify",
	})
}
