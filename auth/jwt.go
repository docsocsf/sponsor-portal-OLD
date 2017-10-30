package auth

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/egnwd/roles"
)

const expiryTime = 3 * time.Hour

type Claims struct {
	jwt.StandardClaims
	UserIdentifier
}

const jtiKey = "jti"

func verifyId(jti string, cmp string, required bool) bool {
	if jti == "" {
		return !required
	}
	return subtle.ConstantTimeCompare([]byte(jti), []byte(cmp)) != 0
}

func GetToken(onetime bool) http.Handler {
	authEnvConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := User(r)
		if user == nil {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
		claims := Claims{
			jwt.StandardClaims{
				Issuer:    authEnvConfig.JwtIssuer,
				ExpiresAt: getExpiry(),
			},
			*user,
		}

		if onetime {
			session, err := cookieJar.Get(r, sessionKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			jti, err := generateAndStore(jtiKey, session, w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			claims.Id = jti
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		json, err := token.SignedString(authEnvConfig.JwtSecret)
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

func RequireJWT(inner http.Handler, auth Auth, validRoles ...string) http.Handler {
	return jwtHandler(inner, auth, false, validRoles...)
}

func RequireOnetimeJWT(inner http.Handler, auth Auth, validRoles ...string) http.Handler {
	return jwtHandler(inner, auth, true, validRoles...)
}

func jwtHandler(inner http.Handler, auth Auth, onetime bool, validRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractToken(r, onetime)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid access", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok {
			if !roles.HasRole(r, &claims.UserIdentifier, validRoles...) {
				http.Error(w, "Invalid access", http.StatusUnauthorized)
				return
			}

			if err = checkJti(claims, onetime, auth, w, r); err != nil {
				http.Error(w, "Invalid access", http.StatusUnauthorized)
				return
			}

			rWithUser := setRequestUser(r, &claims.UserIdentifier)
			inner.ServeHTTP(w, rWithUser)
		} else {
			http.Error(w, "No claims present", http.StatusBadRequest)
		}
	})
}

func extractToken(r *http.Request, onetime bool) (*jwt.Token, error) {
	authEnvConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal(err)
	}

	var tokenStr string
	if onetime {
		tokenStr = r.FormValue("token")
	} else {
		tokenStr = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	}
	return jwt.ParseWithClaims(tokenStr, new(Claims), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return authEnvConfig.JwtSecret, nil
	})
}

func extractJti(auth Auth, w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := auth.session(r, sessionKey)
	if err != nil {
		return "", err
	}

	return getAndDelete(jtiKey, session, w, r)
}

func checkJti(claims *Claims, onetime bool, auth Auth, w http.ResponseWriter, r *http.Request) error {
	jti := claims.Id

	cmp, err := extractJti(auth, w, r)
	if err != nil {
		return err
	}

	if valid := verifyId(jti, cmp, onetime); !valid {
		return fmt.Errorf("JTI token not valid")
	}

	return nil
}
