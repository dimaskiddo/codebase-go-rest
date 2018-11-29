package utils

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dimaskiddo/frame-go/routers"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Token  string `json:"token"`
}

var JWTSigningKey string

func GetJWTToken(data interface{}) (string, error) {
	// Convert Signing Key in Byte Format
	signingKey := []byte(JWTSigningKey)

	// Create JWT Token With HS256 Method and Set Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data":   data.(string),
		"expire": strconv.FormatInt(time.Now().Add(time.Hour*24).Unix(), 10),
	})

	// Generate JWT Token String With Signing Key
	tokenString, err := token.SignedString(signingKey)

	// Return The JWT Token String and Error
	return tokenString, err
}

func GetJWTClaims(data string) (jwt.MapClaims, error) {
	// Convert Signing Key in Byte Format
	signingKey := []byte(JWTSigningKey)

	// Parse JWT Token, If Token is Not Valid Then Return The Signing Key
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	// If Error Found Then Return Empty Claims and The Error
	if err != nil {
		return nil, err
	}

	// Get The Claims
	claims := token.Claims.(jwt.MapClaims)

	// Return The Claims and Error
	return claims, err
}

// Function JWT Authentication as Middleware
func AuthJWT(nextHandlerFunc http.HandlerFunc) http.Handler {
	// Return Next HTTP Handler Function, If Authentication is Valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get JWT Token From HTTP Header Authorization
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if len(token) == 0 {
			routers.ResponseUnauthorized(w)
			return
		}

		// Get Claims From JWT Token
		claims, err := GetJWTClaims(token)
		if err != nil {
			routers.ResponseUnauthorized(w)
			return
		}

		// Set Extracted Claims to HTTP Header
		r.Header.Set("JWT-Data", claims["data"].(string))
		r.Header.Set("JWT-Expire", claims["expire"].(string))

		// Call Next Handler Function With Current Request
		nextHandlerFunc(w, r)
	})
}
