package utils

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ExtractClaims(secret, tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid JWT Token")
}

func TokenVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		parts := strings.Split(token, " ")
		if token == "" {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		claims, err := ExtractClaims(EnvConfigs.SecretApp, parts[1])
		if err != nil {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}
		data := claims["Data"]
		result := map[string]interface{}{}
		encoded, _ := json.Marshal(data)
		json.Unmarshal(encoded, &result)
		var ctx context.Context
		for key, val := range result {
			// c.Set(key, val)
			ctx = context.WithValue(r.Context(), key, val)
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
