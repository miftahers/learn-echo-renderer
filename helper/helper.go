package helper

import (
	config "learn-echo-renderer/Config"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(userID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Encoder
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//return token + error message
	return token.SignedString([]byte(config.Cfg.TokenSecret))
}
