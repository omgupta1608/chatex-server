package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/firebase"
	"github.com/omgupta1608/chatex/server/pkg/middleware/jwt"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

func UserVerificationRouteHandler(c *gin.Context) {
	// parse request body
	var body struct {
		Uid              string `json:"uid"`
		VerificationCode string `json:"verification_code"`
	}
	if err := c.BindJSON(&body); err != nil {
		exception.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	// TODO: verify verification code from redis cache
	isVerified := false
	var user types.User
	for _, u := range UserList {
		if u.Uid == body.Uid {
			isVerified = true
			user = u
			break
		}
	}
	if !isVerified {
		exception.SendError(c, http.StatusUnauthorized, errors.New("Invalid credentials or verification timed out"))
		return
	}

	// write user to database
	_, err := firebase.Client.Collection("Users").Doc(user.Uid).Set(firebase.Ctx, user)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Database Error"))
		return
	}

	// generate jwt token
	jwtTokenString, err := jwt.GenerateTokenString(user.Uid, user.Email)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Cannot generate jwt token"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    jwtTokenString,
		"message": "Login Success",
		"error":   nil,
	})
}
