package UserServices

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	Username   string `json:"username"`
	Authorized bool   `json:"authorized"`
	jwt.RegisteredClaims
}

const invalidToken, unauthorized, tokenExpired = "invalid_auth_token", "unauthorized", "token_expired"

func GenerateJWT(username string, deviceAppKey string) (string, error) {

	claims := CustomClaims{
		username,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    GetEnvValue("APP_MODE"),
			Subject:   "app_user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	privateKey := convertStringToKey(deviceAppKey)
	tokenString, err := token.SignedString(privateKey)
	fmt.Println("token: ", tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(authToken string, deviceAppKey string) (bool, error) {

	parsedToken, err := parseToken(authToken, deviceAppKey)

	if err != nil {
		return false, errors.New(invalidToken + " - it2 " + err.Error())
	}

	claims, ok := parsedToken.Claims.(*CustomClaims)
	if ok && parsedToken.Valid {
		authorized := claims.Authorized
		expiration := claims.RegisteredClaims.ExpiresAt.Time

		if !authorized {
			return false, errors.New(unauthorized + " - it3")
		}

		expired := time.Now().After(expiration)

		if expired {
			return false, errors.New(tokenExpired + " - it4")
		}

	} else {
		return false, errors.New(invalidToken + " - it5" + err.Error())
	}

	return true, nil
}

func convertStringToKey(deviceAppKey string) []byte {
	key := GetEnvValue("APP_KEY") + "-" + deviceAppKey
	privateKey := []byte(key)
	return privateKey
}

func parseToken(authToken string, deviceAppKey string) (*jwt.Token, error) {
	parsedToken, err := jwt.ParseWithClaims(authToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(invalidToken + " - it1")
		}
		secret := convertStringToKey(deviceAppKey)
		return secret, nil
	})

	return parsedToken, err
}

func GetTokenClaims(context *gin.Context) (*CustomClaims, bool) {
	authToken := context.GetHeader("auth-token")
	deviceAppKey := context.GetHeader("app-key")
	parsedToken, _ := parseToken(authToken, deviceAppKey)
	claims, ok := parsedToken.Claims.(*CustomClaims)

	return claims, ok
}
