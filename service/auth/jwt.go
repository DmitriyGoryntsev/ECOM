package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/GDA35/ECOM/types"
	"github.com/GDA35/ECOM/utils"
	"github.com/golang-jwt/jwt"
)

type ContextKey string

const UserKey ContextKey = "userID"

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

func WithJWTAuth(next http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get token from the user request
		tokenString := getTokenFromRequest(r)

		//validate token
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("invalid token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		//if is we need to fetch user from the database
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		//add user to the context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		//call next handler
		next(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth == "" {
		return ""
	}

	return tokenAuth
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
