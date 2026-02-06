package middleware

import (
	"context"
	"net/http"

	"Concurrent_Task_Management_System/internal/repositories"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func AuthMiddleware(userRepo repositories.UserRepository) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 🔓 Allow user creation without auth
			if r.Method == http.MethodPost && r.URL.Path == "/users" {
				next.ServeHTTP(w, r)
				return
			}

			userID := r.Header.Get("X-User-Id")
			if userID == "" {
				http.Error(w, "X-User-Id header missing", http.StatusUnauthorized)
				return
			}

			objID, err := primitive.ObjectIDFromHex(userID)
			if err != nil {
				http.Error(w, "invalid X-User-Id", http.StatusUnauthorized)
				return
			}

			user, err := userRepo.FindByID(r.Context(), objID)
			if err != nil {
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "currentUser", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
