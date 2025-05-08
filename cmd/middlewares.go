package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/Auxesia23/todo_list/internal/utils"
	"github.com/golang-jwt/jwt"
)

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.VerifyJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userEmail, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Invalid email in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userEmail", userEmail)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
