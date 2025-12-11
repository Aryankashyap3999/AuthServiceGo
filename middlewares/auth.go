package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	dbConfig "AuthInGo/config/db"
	repo "AuthInGo/db/repositories"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func (token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: " + err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		fmt.Println("userId:", userId, " email:", email)
		
		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}	

		fmt.Printf("Authenticated user ID: %v, Email: %s\n", int64(userId), email)

		ctx := context.WithValue(r.Context(), "userId", strconv.FormatFloat(userId, 'f', 0, 64))	
		ctx = context.WithValue(ctx, "email", email)	

		next.ServeHTTP(w, r.WithContext(ctx))
		
	})
}

func RequireAllRoles(roles ...string) func (http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			
			userIdStr := r.Context().Value("userId").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)

			if err != nil {
				http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDB()

			if dbErr != nil {
				http.Error(w, "Database connection error", http.StatusInternalServerError)
				return
			}
			
			urr := repo.NewUserRoleRepository(dbConn)

			hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId, roles)

			if hasAllRolesErr != nil {
				http.Error(w, "Error checking roles: "+hasAllRolesErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllRoles {
				http.Error(w, "Forbidden: missing required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User has all required roles:", roles)
			
			next.ServeHTTP(w, r)
		})
	}
}

func RequireAnyRole(roles ...string) func (http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			
			userIdStr := r.Context().Value("userId").(string)
			userId, err := strconv.ParseInt(userIdStr, 10, 64)

			if err != nil {
				http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
				return
			}

			dbConn, dbErr := dbConfig.SetupDB()

			if dbErr != nil {
				http.Error(w, "Database connection error", http.StatusInternalServerError)
				return
			}
			
			urr := repo.NewUserRoleRepository(dbConn)

			hasAnyRoles, hasAnyRolesErr := urr.HasAnyRole(userId, roles)

			if hasAnyRolesErr != nil {
				http.Error(w, "Error checking roles: "+hasAnyRolesErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAnyRoles {
				http.Error(w, "Forbidden: missing required roles", http.StatusForbidden)
				return
			}

			fmt.Println("User has any required roles:", roles)
			
			next.ServeHTTP(w, r)
		})
	}
}
