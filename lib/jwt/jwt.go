package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func NewToken(userID uint64, appID, appSecret string, expTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = userID
	claims["exp"] = expTime
	claims["app_id"] = appID

	tokenString, err := token.SignedString([]byte(appSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString, secret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return err
		//log.Fatal(err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return err
	}

	return nil
}
