package jwt

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	// signing algorithm. Required
	// possible values HS256, HS384, HS512, RS256, RS384 or RS512
	// defaults to HS256
	SigningMethodStr string

	// Secret key used for signing. Required.
	SecretKey []byte

	// key of cookie containing jwt token. Required
	CookieName string

	// Number of seconds until the cookie expires
	CookieMaxAge int

	// Cookie is only sent to the server when a request is made with the https: scheme (except on localhost)
	SecureCookie bool

	// Forbids JavaScript from accessing the cookie
	HttpOnlyCookie bool

	// A path that must exist in the requested URL for cookie to be sent
	CookiePath string

	// Host to which the cookie will be sent
	CookieDomain string

	// Signing method
	signingMethod jwt.SigningMethod
}

type AuthCustomClaims struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// create and validate a new JWTService
func New(service *JWTService) *JWTService {
	// required fields
	if service.SigningMethodStr == "" {
		service.SigningMethodStr = "HS256"
	}

	if service.SecretKey == nil {
		panic(errors.New("Secret key must be passed"))
	}

	if service.CookieName == "" {
		panic(errors.New("Cookie name must be passed"))
	}

	service.signingMethod = jwt.GetSigningMethod(service.SigningMethodStr)

	return service
}

// generate jwt using data provided as payload
func (service *JWTService) GenerateTokenString(uid, email string) (string, error) {
	claim := AuthCustomClaims{uid, email, jwt.StandardClaims{}}
	token := jwt.NewWithClaims(service.signingMethod, claim)

	// sign the token using secret key
	signedToken, err := token.SignedString(service.SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// validate token and extract claims
func (service *JWTService) ParseToken(tokenString string) (*AuthCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AuthCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if service.signingMethod.Alg() != t.Method.Alg() {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return service.SecretKey, nil
		},
	)

	// extract claims
	claims, ok := token.Claims.(*AuthCustomClaims)
	if err == nil && ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid Token, parse error")
}
