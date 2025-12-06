package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/farizziezhi/layered/internal/domain"
	"github.com/farizziezhi/layered/pkg/helper"
)

var jwtKey = "alamak_kasus_ni"
func AuthMiddleware(next http.Handler) http.Handler {
     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         authHeader := r.Header.Get("Authorization")
         if authHeader == "" {
              helper.SendJSON(w, http.StatusUnauthorized, domain.Response{Message: "Token tidak ditemukan"})
              return
         }

         tokenString := strings.TrimPrefix(authHeader, "Bearer ")
         if tokenString == authHeader { 
             helper.SendJSON(w, http.StatusUnauthorized, domain.Response{Message: "Format token salah"})
             return
         }

         claims := &domain.Claims{}
         token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
              return []byte(jwtKey), nil
         })
         if err != nil || !token.Valid {
              helper.SendJSON(w, http.StatusUnauthorized, domain.Response{Message: "Token tidak valid"})
              return
         }
          // Lanjut ke handler berikutnya
         next.ServeHTTP(w, r)
     })
}