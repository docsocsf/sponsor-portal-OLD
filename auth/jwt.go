package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/docsocsf/sponsor-portal/config"
)

const expiryTime = 3 * time.Hour

type Claims struct {
	jwt.StandardClaims

	User UserIdentifier `json:"user,omitempty"`
}

var (
	issuer string
	secret []byte
)

func init() {
	authEnvConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal(err, "Make jwt service")
	}

	issuer = authEnvConfig.JwtIssuer
	secret = []byte(authEnvConfig.JwtSecret)
}

func GetToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := Claims{
			jwt.StandardClaims{
				Issuer:    issuer,
				ExpiresAt: getExpiry(),
			},
			User(r),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		json, err := token.SignedString(secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(json))
	})
}

func getExpiry() int64 {
	return time.Now().Add(expiryTime).Unix()
}

func RequireJWT(auth Auth, inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, new(Claims), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return secret, nil
		})
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			rWithUser := setRequestUser(r, claims.User)
			inner.ServeHTTP(w, rWithUser)
		}
	})
}
