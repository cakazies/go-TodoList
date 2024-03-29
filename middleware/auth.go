package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/local/TaskListGo/models"
	util "github.com/local/TaskListGo/util"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuthPath := []string{"/api/user/register"}
		noAuthPath = append(noAuthPath, "/api/user/login")
		requestPath := r.URL.Path
		for _, path := range noAuthPath {
			log.Println(requestPath, " dan ", path)
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = util.MetaMsg(false, "Token is not present")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		headerAuthorizationString := strings.Split(tokenHeader, " ")
		if len(headerAuthorizationString) != 2 {
			response = util.MetaMsg(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		tk := &models.Token{}
		tokenValue := headerAuthorizationString[1]
		token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwt_secret")), nil
		})

		if err != nil {
			response = util.MetaMsg(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		if !token.Valid {
			response = util.MetaMsg(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Respond(w, response)
			return
		}

		fmt.Sprintf("User Id is %s", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
