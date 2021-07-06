package handlers

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/types"
	"github.com/omgupta1608/chatex/server/pkg/validation"
	"github.com/rs/xid"
)

// ! not for production
var UserList []types.User

func RegisterRouteHandler(c *gin.Context) {
	// user data from request body
	var reqBody struct {
		Name     string `json:"name" validate:"required,min=3,max=15"`
		About    string `json:"about" validate:"required,min=3,max=40"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=40"`
	}

	err := c.BindJSON(&reqBody)
	if err != nil {
		exception.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	// TODO: sanitize client input
	errFields, invalidValidationError := validation.ValidateReqData(&reqBody)
	if invalidValidationError != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("InvalidValidationError"))
		return
	}
	if len(errFields) != 0 {
		exception.SendValidationError(c, errFields)
		return
	}

	// initialize/generate user data
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Could not hash the password"))
		return
	}

	// create a new user
	newUser := types.User{
		Uid:        xid.New().String(),
		Name:       reqBody.Name,
		About:      reqBody.About,
		Email:      reqBody.Email,
		Password:   string(hashedPassword),
		ProfilePic: "", // TODO: add a default profile Pic,
	}
	if newUser.About == "" {
		newUser.About = "Hello there!! I'm using ChatX"
	}

	// TODO: store this user in redis and wait for verification
	UserList = append(UserList, newUser)

	// TODO: create a verification token & send through email
	// TODO: store verify token in redis

	c.JSON(http.StatusCreated, gin.H{
		"data":               newUser.Uid,
		"message":            "User created",
		"error":              nil,
		"verification-route": c.FullPath() + "/verify",
	})
}
