package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/joaolima7/maconaria_back-end/internal/infra/web/auth"
	"github.com/joaolima7/maconaria_back-end/internal/infra/web/response"
)

type contextKey string

const UserContextKey contextKey = "user"

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Unauthorized(w, "Token de autenticação não fornecido!", nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer")
		if tokenString == authHeader {
			response.Unauthorized(w, "Formasto de Token inválido!", nil)
		}

		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(w, "Token inválido ou expirado!", err)
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserContextKey).(*auth.JWTClaims)
		if !ok {
			response.Unauthorized(w, "Token inválido!", nil)
			return
		}

		if !claims.IsAdmin {
			response.Forbidden(w, "Acesso negado: permissões de administrador necessárias!", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) (*auth.JWTClaims, bool) {
	claims, ok := ctx.Value(UserContextKey).(*auth.JWTClaims)
	return claims, ok
}
