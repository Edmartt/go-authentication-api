package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/time/rate"

	"github.com/Edmartt/go-authentication-api/internal/users/models"
	"github.com/Edmartt/go-authentication-api/pkg/jwt"
)


func SetJSONResponse(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		
		handler.ServeHTTP(w, r)
	}
}

func ValidateRequestBody(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
	rBody, rError := io.ReadAll(r.Body)

	user := models.User{}

	if rError != nil{
		log.Println(rError.Error())
	}

	json.Unmarshal(rBody, &user)

	r.Body = ioutil.NopCloser(bytes.NewBuffer(rBody))

	if len(user.Username) < 5{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The username must be 5 characters min.")
		return
	}

	if len(user.Password) < 8{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The password must be 8 characters min.")
		return
	}

	handler(w, r)
	}	
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		sentToken := r.Header.Get("Authorization")

		if sentToken == ""{
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("No Authorization Header")
			return
		}

		tokenFromRequest := strings.Split(sentToken, "Bearer")

		if len(tokenFromRequest) == 2{
			sentToken = strings.TrimSpace(tokenFromRequest[1])
		}else{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Incorrect Format")
			return
		}

		jwtWrapper := jwt.JWTWrapper{
		}

		claims, err := jwtWrapper.ValidateToken(sentToken)

		if err != nil{
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Attribute)
		handler(w, r.WithContext(ctx))
	}
}



func LimitRequest(next http.Handler) http.Handler{
	limit := rate.NewLimiter(0.3, 3)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if limit.Allow() == false{
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
