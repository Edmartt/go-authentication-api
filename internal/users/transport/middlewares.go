package transport

import (
	"bytes"
	"encoding/json"
	"github.com/Edmartt/go-authentication-api/internal/users/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
