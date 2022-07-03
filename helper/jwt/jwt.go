package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("AllYourBase")
	issure    = "Pien Corporation"
)

type Claims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

func Generate(uuid string) (string, error) {
	claims := Claims{
		UUID: uuid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issure,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)
	return ss, err
}

func Validate(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid Signature")
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ExtractClaims(tokenString string, key string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid Signature")
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims[key])
	} else {
		return err.Error()
	}
}
