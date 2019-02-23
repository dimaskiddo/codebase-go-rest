package service

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// FormatGetJWT Struct
type FormatGetJWT struct {
	Status  bool              `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

// jwtKeysConfig Struct
type jwtKeysConfig struct {
	Private []byte
	Public  []byte
}

// jwtKeysConfig Variable
var jwtKeysCfg jwtKeysConfig

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
func GetJWTToken(payload interface{}) (string, error) {
	// Convert Signing Key in Byte Format
	signingKey, err := jwt.ParseRSAPrivateKeyFromPEM(jwtKeysCfg.Private)
	if err != nil {
		return "", err
	}

	// Create JWT Token With RS256 Method And Set JWT Claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"data": payload.(string),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Generate JWT Token String With Signing Key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	// Return The JWT Token String and Error
	return tokenString, nil
}

// JWTClaims Function to Get JWT Claims Information
func jwtClaims(data string) (jwt.MapClaims, error) {
	// Convert Signing Key in Byte Format
	signingKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtKeysCfg.Public)
	if err != nil {
		return nil, err
	}

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
