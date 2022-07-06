package services

import (
	"fmt"
	"time"

	"github.com/abelgalef/course-reg/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenService interface {
	GenerateToken(models.User) (string, error)
	ValidateToken(string) (interface{}, bool)
}

type tokenService struct {
	authKey string
	issuer  string
}

type claims struct {
	User models.User
	jwt.StandardClaims
}

func NewJWTService() JWTTokenService {
	return &tokenService{"my_secrte_key", "Course-Registration_(Defense-Enginering-University)"}
}

func (ts *tokenService) GenerateToken(user models.User) (string, error) {
	claims := &claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8760).Unix(), // Valid for a year
			Issuer:    ts.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := token.SignedString(([]byte(ts.authKey)))
	if err != nil {
		return "", err
	}

	return tk, nil
}

func (ts *tokenService) ValidateToken(tk string) (interface{}, bool) {
	token, err := jwt.Parse(tk, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(ts.authKey), nil
	})
	if err != nil {
		return nil, false
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return claims["User"], true
	}
	fmt.Println("herrreeee")
	return nil, false
}
