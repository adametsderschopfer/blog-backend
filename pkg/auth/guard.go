package auth

import (
	"app/internal/service"
	"app/pkg/domain/entity"
	"net/http"
)

func CreateGuardMiddleware(authService service.Auth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, authId, err := findAuthData(r)
			if err != nil || accessToken == "" || authId == "" {
				http.Error(w, "Required auth fields is missing", http.StatusUnauthorized)
				return
			}

			err = authService.VerifyToken(authId, accessToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func findAuthData(r *http.Request) (string, string, error) {
	accessTokenHeader := r.Header.Get(entity.AccessTokenHeaderFieldName)
	authIdHeader := r.Header.Get(entity.AccessTokenHeaderFieldName)

	if accessTokenHeader != "" && authIdHeader != "" {
		return accessTokenHeader, authIdHeader, nil
	}

	accessTokenCookie, _ := r.Cookie(entity.AccessTokenFieldName)
	if err := accessTokenCookie.Valid(); err != nil {
		return "", "", err
	}

	authIdCookie, _ := r.Cookie(entity.AuthIdFieldName)
	if err := authIdCookie.Valid(); err != nil {
		return "", "", err
	}

	return accessTokenCookie.Value, authIdCookie.Value, nil
}
