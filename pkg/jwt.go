package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewToken(uid, role string) *claims {
	return &claims{
		Id:   uid,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "backendgolang",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
		},
	}
}

func (c *claims) Generate() (string, error) {
	secrets := "kyt"
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString([]byte(secrets))
}

func VerifyToken(token string) (*claims, error) {
	secrets := "kyt"
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets), nil
	})

	if err != nil {
		return nil, err
	}

	claimData := data.Claims.(*claims)
	return claimData, nil

}
