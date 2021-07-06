package handlers

import (
	"errors"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/chatex/server/pkg/exception"
	"github.com/omgupta1608/chatex/server/pkg/firebase"
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

// Handler function for Edit User Profile route
// @param {string} uid - Id of the User
// /api/v/user/edit/:uid
func EditUserProfile(c *gin.Context) {
	// Uid, Email Cannot be Changed
	var data struct {
		Name       string `json:"name,string"`
		About      string `json:"about,string"`
		ProfilePic string `json:"profile_pic,string"`
	}

	if err := c.BindJSON(&data); err != nil {
		exception.SendError(c, http.StatusBadRequest, err)
		return
	}

	ID, _ := c.Params.Get("uid")

	_, err := firebase.Client.Collection("Users").Doc(ID).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, http.StatusNotFound, err)
		return
	}

	_, err = firebase.Client.Collection("Users").Doc(ID).Set(firebase.Ctx, data, firestore.MergeAll)

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
	var data struct {
		OldPassword string `json:"old_password,string"`
		NewPassword string `json:"new_password,string"`
	}

	if err := c.BindJSON(&data); err != nil {
		exception.SendError(c, http.StatusBadRequest, err)
		return
	}

	ID, _ := c.Params.Get("uid")

	dbsnapshot, err := firebase.Client.Collection("Users").Doc(ID).Get(firebase.Ctx)

	if err != nil {
		exception.SendError(c, http.StatusNotFound, err)
		return
	}

	user := dbsnapshot.Data()

	if data.OldPassword == user["Password"] {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 10)
		if err != nil {
			exception.SendError(c, http.StatusInternalServerError, errors.New("Could not hash the password"))
			return
		}

		_, err = firebase.Client.Collection("Users").Doc(ID).Set(firebase.Ctx, map[string]interface{}{
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
	} else {
		exception.SendError(c, http.StatusUnauthorized, errors.New("Invalid Password"))
	}
}

// Handler function for Delete User By Id route
// @param {string} uid - Id of the User
// /api/v/user/delete-account/:uid
func DeleteUserById(c *gin.Context) {
	ID, _ := c.Params.Get("uid")

	_, err := firebase.Client.Collection("Users").Doc(ID).Delete(firebase.Ctx)
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
