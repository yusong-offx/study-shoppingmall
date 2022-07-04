package utils

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var serete_key = []byte(`secrete_key`)

type AuthClaims struct {
	Id        string
	UserType  string
	ExpiredAt int64
	Uuid      uuid.UUID
	jwt.Claims
}

func (a AuthClaims) Valid() error {
	if a.ExpiredAt <= time.Now().Unix() {
		return errors.New("time expired")
	}
	// Refresh token 구현
	return nil
}

func GenerateToken(authclaims AuthClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authclaims)
	signedtoken, err := token.SignedString(serete_key)
	if err != nil {
		return "", err
	}
	return signedtoken, nil
}

func VerfyJWT(token string) error {
	_, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("diff jwt method")
		}
		return serete_key, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func JWTmiddleware(c *fiber.Ctx) error {
	err := VerfyJWT(c.GetReqHeaders()["Authorization"])
	if err != nil {
		return ErrorReqeustJSON(err, 403, c)
	}
	return c.Next()
}
