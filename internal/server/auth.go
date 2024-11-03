package server

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	JWTKey []byte
}

func NewAuth(key string) *Auth {
	return &Auth{
		JWTKey: []byte(key),
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (u *Auth) JwtGenerator(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.JWTKey)
}

func (u *Auth) JwtValidator(tokenStr string) (*Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return u.JWTKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return &claims, nil
}
