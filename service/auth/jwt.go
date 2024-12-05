package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userID int) (string, error) {

	// not a best practice
	JWTExpirationInSeconds := 3600 * 24 * 7
	expiration := time.Second * time.Duration(JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userID": strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix()})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
