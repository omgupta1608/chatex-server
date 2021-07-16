package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/firebase"
	"github.com/omgupta1608/chatex/server/pkg/middleware/jwt"
	"github.com/omgupta1608/chatex/server/pkg/types"
	"github.com/omgupta1608/chatex/server/pkg/validation"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

func LoginRouteHandler(c *gin.Context) {
	var reqData types.LoginReqData
	if err := c.BindJSON(&reqData); err != nil {
		exception.SendError(c, http.StatusBadRequest, errors.New("Bad JSON format"))
		return
	}

	// Never trust client side data
	// TODO: sanitize client input
	errFields, invalidValidationError := validation.ValidateReqData(&reqData)
	if invalidValidationError != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("InvalidValidationError"))
		return
	}
	if len(errFields) != 0 {
		exception.SendValidationError(c, errFields)
		return
	}

	user, httpStatusCode, err := getUserByEmail(c, reqData.Email)
	if err != nil {
		exception.SendError(c, httpStatusCode, err)
		return
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqData.Password))
	if err != nil {
		exception.SendError(c, http.StatusUnauthorized, errors.New("Email Or Password is invalid"))
		return
	}

	// create jwt token
	jwtTokenString, err := jwt.GenerateTokenString(user.Uid, user.Email)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Cannot generate jwt token"))
		return
	}

	// remove password from response
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"jwt":     jwtTokenString,
		"user":    user,
		"error":   nil,
	})
}

// search firestore for user with given email
func getUserByEmail(c *gin.Context, email string) (user *types.User, httpStatusCode int, err error) {
	iter := firebase.Client.Collection("Users").Where("Email", "==", email).Documents(firebase.Ctx)
	doc, err := iter.Next()

	// user not found
	if err == iterator.Done {
		err = errors.New("Email Or Password is invalid")
		return nil, http.StatusUnauthorized, err
	}
	if err != nil {
		err = errors.New("Database error")
		return nil, http.StatusInternalServerError, err
	}

	// unmarshal user
	err = doc.DataTo(&user)
	if err != nil {
		err = errors.New("Database error")
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
