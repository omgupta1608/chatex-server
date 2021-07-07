package handlers

import (
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/firebase"
	"github.com/omgupta1608/chatex/server/pkg/types"
	"github.com/omgupta1608/chatex/server/pkg/validation"
	"golang.org/x/crypto/bcrypt"
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

	// validate uid
	if validId := validation.ValidateID(uid); !validId {
		exception.SendError(c, 403, errors.New("invalid user id"))
		return
	}

	// find user with uid in DB
	dbsnapshot, err := firebase.Client.Collection("Users").Doc(uid).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, 500, nil)
		return
	}

	user := dbsnapshot.Data()

	// remove the password field
	delete(user, "Password")

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "",
		"error":   nil,
	})
}

// Handler function for Edit User Profile route
// @param {string} uid - Id of the User
// /api/v/user/edit/:uid
func EditUserProfile(c *gin.Context) {
	// Uid, Email Cannot be Changed
	var data types.EditUserProfileReqData

	if err := c.BindJSON(&data); err != nil {
		exception.SendError(c, http.StatusBadRequest, err)
		return
	}

	// validate user input
	errFields, invalidValidationError := validation.ValidateReqData(&data)
	if invalidValidationError != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("InvalidValidationError"))
		return
	}
	if len(errFields) != 0 {
		exception.SendValidationError(c, errFields)
		return
	}

	// get id from params
	uid, _ := c.Params.Get("uid")

	// validate uid
	if validId := validation.ValidateID(uid); !validId {
		exception.SendError(c, 403, errors.New("invalid user id"))
		return
	}

	// check if user exists
	_, err := firebase.Client.Collection("Users").Doc(uid).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, http.StatusNotFound, err)
		return
	}

	// update user in DB
	_, err = firebase.Client.Collection("Users").Doc(uid).Set(firebase.Ctx, data, firestore.MergeAll)

	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Profile Edited Successfully",
		"error":   nil,
	})
}

// Handler function for Change User Password route
// @param {string} uid - Id of the User
// /api/v/user/change-password/:uid
func ChangePassword(c *gin.Context) {
	var data types.ChangePasswordReqData

	if err := c.BindJSON(&data); err != nil {
		exception.SendError(c, http.StatusBadRequest, err)
		return
	}

	// validate user input
	errFields, invalidValidationError := validation.ValidateReqData(&data)
	if invalidValidationError != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("InvalidValidationError"))
		return
	}
	if len(errFields) != 0 {
		exception.SendValidationError(c, errFields)
		return
	}

	// get id from params
	uid, _ := c.Params.Get("uid")

	// validate uid
	if validId := validation.ValidateID(uid); !validId {
		exception.SendError(c, 403, errors.New("invalid user id"))
		return
	}

	// check if user exists
	dbsnapshot, err := firebase.Client.Collection("Users").Doc(uid).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, http.StatusNotFound, err)
		return
	}

	user := dbsnapshot.Data()

	// verify if the password provided is correct
	err = bcrypt.CompareHashAndPassword([]byte(user["Password"].(string)), []byte(data.OldPassword))
	if err != nil {
		exception.SendError(c, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	// encrypt the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 10)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, errors.New("Could not hash the password"))
		return
	}

	// update password in DB
	_, err = firebase.Client.Collection("Users").Doc(uid).Set(firebase.Ctx, map[string]interface{}{
		"Password": hashedPassword,
	}, firestore.MergeAll)

	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "Password changed Successfully!",
		"error":   nil,
	})

}

// Handler function for Delete User By Id route
// @param {string} uid - Id of the User
// /api/v/user/delete-account/:uid
func DeleteUserById(c *gin.Context) {
	// get id from params
	uid, _ := c.Params.Get("uid")

	// validate uid
	if validId := validation.ValidateID(uid); !validId {
		exception.SendError(c, 403, errors.New("invalid user id"))
		return
	}

	// delete user from DB
	_, err := firebase.Client.Collection("Users").Doc(uid).Delete(firebase.Ctx)
	if err != nil {
		exception.SendError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "Account Delete Successfully",
		"error":   nil,
	})
}
