package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
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
			response.Error(w, apperrors.NewUnauthorizedError("Token de autenticação não fornecido"))
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			response.Error(w, apperrors.NewUnauthorizedError("Formato de token inválido. Use: 'Authorization: Bearer <token>'"))
			return
		}

		tokenString := strings.TrimSpace(authHeader[len(bearerPrefix):])
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {

			response.Error(w, apperrors.NewUnauthorizedError("Token inválido ou expirado"))
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserContextKey).(*auth.JWTClaims)
		if !ok {
			response.Error(w, apperrors.NewUnauthorizedError("Token inválido"))
			return
		}

		if !claims.IsAdmin {
			response.Error(w, apperrors.NewForbiddenError("Acesso negado: permissões de administrador necessárias"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetUserFromContext(ctx context.Context) (*auth.JWTClaims, bool) {
	claims, ok := ctx.Value(UserContextKey).(*auth.JWTClaims)
	return claims, ok
}
