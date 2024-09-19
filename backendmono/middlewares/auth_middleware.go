package middlewares // Baris ini harus ada di paling atas file

import (
    "context"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gorilla/mux"
    "net/http"
    "strings"
    "log"
)

func JWTAuthMiddleware(jwtSecret string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            log.Println("JWT Middleware invoked")

            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                log.Println("Authorization header missing")
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }

            parts := strings.Split(authHeader, "Bearer ")
            if len(parts) != 2 {
                log.Println("Invalid Authorization header format")
                http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
                return
            }

            tokenString := parts[1]
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte(jwtSecret), nil
            })

            if err != nil || !token.Valid {
                log.Println("Invalid token")
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok || claims["email"] == nil {
                log.Println("Invalid token claims")
                http.Error(w, "Invalid token claims", http.StatusUnauthorized)
                return
            }

            email := claims["email"].(string)
            log.Println("Email saved in context:", email)

            ctx := context.WithValue(r.Context(), "email", email)
            r = r.WithContext(ctx)

            next.ServeHTTP(w, r)
        })
    }
}
