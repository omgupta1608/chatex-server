package jwt

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/omgupta1608/chatex/server/cmd/config"
	"github.com/omgupta1608/chatex/server/pkg/types"
)

type AuthCustomClaims struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var signingMethod = jwt.SigningMethodHS256
var secretKey = config.GetJwtSecret()

// generate jwt using data provided as payload
func GenerateTokenString(uid, email string) (string, error) {
	claim := AuthCustomClaims{uid, email, jwt.StandardClaims{}}
	token := jwt.NewWithClaims(signingMethod, claim)

	// sign the token using secret key
	return token.SignedString(secretKey)
}

// validate token and extract claims
func ParseToken(tokenString string) (*types.User, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AuthCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if signingMethod.Alg() != t.Method.Alg() {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return secretKey, nil
		},
	)

	// extract claims
	claims, ok := token.Claims.(*AuthCustomClaims)
	if err == nil && ok && token.Valid {
		user := types.User{Uid: claims.Uid, Email: claims.Email}
		return &user, nil
	}

	return nil, errors.New("Invalid Token, parse error")
}
