package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	Email    string `json:"email"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	Sub      string `json:"sub"`
	Iat      int64  `json:"iat"`
}

type JwtService interface {
	Sign(email CustomJwtClaims) (string, error)

	Verify(token string) bool
}

type jwtService struct{}

func NewJwtService() JwtService {
	return &jwtService{}
}

const (
	mySecret = "my-secret-key"
)

func (s *jwtService) Sign(claims CustomJwtClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    claims.Email,
		"exp":      claims.Exp,
		"username": claims.Username,
		"iat":      claims.Iat,
		"sub":      claims.Sub,
	})

	return token.SignedString([]byte(mySecret))
}

func (s *jwtService) Verify(receivedToken string) bool {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySecret, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"] == "lenin"
	}

	return false
}
