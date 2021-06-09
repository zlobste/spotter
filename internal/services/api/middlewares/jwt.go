package middlewares

import (
	"github.com/zlobste/spotter/internal/utils"
	"net/http"
)

// JWTMiddleware verifies JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := utils.TokenValid(r)
		if err != nil {
			utils.Respond(w, http.StatusUnauthorized, utils.Message(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	})
}
