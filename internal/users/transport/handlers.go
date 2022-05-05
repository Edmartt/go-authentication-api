package transport

import (
	"encoding/json"
	"io"
	"log"

	"github.com/Edmartt/go-authentication-api/internal/users/data"
	"github.com/Edmartt/go-authentication-api/internal/users/models"
	"github.com/Edmartt/go-authentication-api/pkg/jwt"

	"net/http"

	"github.com/google/uuid"

	"github.com/Edmartt/go-password-hasher/hasher"
)


type Handlers struct {
	userRepo data.UserRepository
}

func(h *Handlers) Login(w http.ResponseWriter, request *http.Request){
	reqBody, requestError := io.ReadAll(request.Body)


	if requestError != nil{
		log.Println(requestError.Error())
	}


	user := &models.User{}

	json.Unmarshal(reqBody, &user)

	searchedUser := h.userRepo.Find(user.Username)


	if searchedUser.Username == user.Username{
		if hasher.CheckHash(searchedUser.Password, user.Password){
			claims := jwt.Claims{}

			newToken := claims.GenerateJWT(user.Username, 5)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(newToken)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Wrong username or password")

		return
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Username or Password Wrong")

	return
}

func (h *Handlers)Signup(w http.ResponseWriter, request *http.Request){

	user := models.User{}
	
	user.Id = uuid.NewString()
	requestError := json.NewDecoder(request.Body).Decode(&user)

	hashedPassword := hasher.ConvertToHash(user.Password)
	user.Password = hashedPassword

	if requestError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	h.userRepo.Create(user)
	w.WriteHeader(http.StatusCreated)
	return
}
