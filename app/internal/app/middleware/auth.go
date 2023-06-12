package middleware

import (
	tokenPkg "app/internal/pkg/auth/token"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func AuthMiddleware(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AuthMiddleware", r.URL.Path)
		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return tokenPkg.SECRET, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if len(roles) > 0 {
			roleMatch := false
			for _, role := range roles {
				if claims["role"] == role {
					roleMatch = true
					break
				}
			}
			if !roleMatch {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
