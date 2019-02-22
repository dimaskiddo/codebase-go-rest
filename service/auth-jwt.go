package service

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTResponse Struct
type JWTResponse struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Token  string `json:"token"`
}

// JWTSigningKey Variable
var jwtSigningKey string

// AuthJWT Function as Midleware for JWT Authorization
func AuthJWT(nextHandlerFunc http.HandlerFunc) http.Handler {
	// Return Next HTTP Handler Function, If Authorization is Valid
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse HTTP Header Authorization
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		// Check HTTP Header Authorization Section
		// Authorization Section Length Should Be 2
		// The First Authorization Section Should Be "Bearer"
		if len(authHeader) != 2 || authHeader[0] != "Bearer" {
			ResponseUnauthorized(w)
			return
		}

		// The Second Authorization Section Should Be The Credentials Payload
		authPayload := authHeader[1]
		if len(authPayload) == 0 {
			ResponseUnauthorized(w)
			return
		}

		// Get Authorization Claims From JWT Token
		authClaims, err := jwtClaims(authPayload)
		if err != nil {
			ResponseUnauthorized(w)
			return
		}

		// Set Extracted Authorization Claims to HTTP Header
		// With Base64 Format
		r.Header.Set("X-JWT-Claims", base64.StdEncoding.EncodeToString([]byte(authClaims["data"].(string))))

		// Call Next Handler Function With Current Request
		nextHandlerFunc(w, r)
	})
}

// GetJWTToken Function to Generate JWT Token
func GetJWTToken(data interface{}) (string, error) {
	// Convert Signing Key in Byte Format
	signingKey := []byte(jwtSigningKey)

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

// JWTClaims Function to Get JWT Claims Information
func jwtClaims(data string) (jwt.MapClaims, error) {
	// Convert Signing Key in Byte Format
	signingKey := []byte(jwtSigningKey)

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
