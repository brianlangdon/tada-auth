package auth

import (
	"context"
	"net/http"

	"github.com/brianlangdon/tada-auth/graph/model"
	"github.com/brianlangdon/tada-auth/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}
			//validate jwt token
			tokenStr := header
			auth := tokenStr[len("Bearer "):]

			username, err := jwt.ParseToken(auth)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and add to context
			user := model.FullUser{Username: username}
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.FullUser {
	raw, _ := ctx.Value(userCtxKey).(*model.FullUser)
	return raw
}
