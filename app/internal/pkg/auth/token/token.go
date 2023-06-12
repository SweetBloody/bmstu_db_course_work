package token

import (
	"app/internal/pkg/models"
	"github.com/golang-jwt/jwt"
	"time"
)

var SECRET = []byte("sseeccrreett")

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": user.Login,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
