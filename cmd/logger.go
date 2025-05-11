package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Auxesia23/todo_list/internal/models"
	"github.com/Auxesia23/todo_list/internal/utils"
	"github.com/golang-jwt/jwt"
)


func(app *application) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Catat waktu mulai
		start := time.Now()

		// Tangkap status code
		rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		// Ambil IP
		ip := r.RemoteAddr
		if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
			ip = strings.Split(forwarded, ",")[0]
		}

		// Ambil email dari JWT jika ada
		var email string
		email = "anonymous"
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
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

			emailClaim, ok := claims["email"].(string)
			if ok {
				email = emailClaim
			}
		}

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(rec, r)

		// Log hasil
		log.Printf("[LOG] %s - %s - %s %s - %d - %s - %v",
			ip,
			email,
			r.Method,
			r.URL.Path,
			rec.statusCode,
			http.StatusText(rec.statusCode),
			time.Since(start),
		)
		
		logsEntry := models.LogEntry{
			IP: ip,
			Email: email,
			Method: r.Method,
			Path: r.URL.Path,
			Status: rec.statusCode,
			Message: http.StatusText(rec.statusCode),
			Duration: time.Since(start).String(),
			Timestamp: time.Now(),
		}
		
		err := app.Logs.Create(context.Background(),logsEntry)
		if err != nil {
			log.Println(err)
		}
		
	}) 
}


// Custom ResponseWriter untuk menangkap status code
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
