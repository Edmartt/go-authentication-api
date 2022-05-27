package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/Edmartt/go-authentication-api/internal/users/data"
	"github.com/Edmartt/go-authentication-api/internal/users/models"
	"github.com/Edmartt/go-authentication-api/pkg/jwt"

	"net/http"

	"github.com/google/uuid"

	"github.com/Edmartt/go-password-hasher/hasher"
)



//Handler struct gives access to user data access layer
type Handlers struct {
	userRepo data.IUserRepository
	user models.User
	logResponse LoginResponse
	sigResponse SignupResponse
	wrapper jwt.JWTWrapper

}

//Login endpoint
func(h *Handlers) Login(w http.ResponseWriter, request *http.Request){
	reqBody, requestError := io.ReadAll(request.Body)

	if requestError != nil{
		log.Println(requestError.Error())
	}

	json.Unmarshal(reqBody, &h.user)

	searchedUser := h.userRepo.Find(h.user.Username)


	if searchedUser.Username == h.user.Username{
		if hasher.CheckHash(searchedUser.Password, h.user.Password){

			newToken, err := h.wrapper.GenerateJWT(h.user.Username, 5)

			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			h.logResponse.Token = newToken
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&h.logResponse)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Wrong username or password")

		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Username or Password Wrong")

	return
}

func (h *Handlers)Signup(w http.ResponseWriter, request *http.Request) {

	h.user.Id = uuid.NewString()
	requestError := json.NewDecoder(request.Body).Decode(&h.user)

	hashedPassword := hasher.ConvertToHash(h.user.Password)
	h.user.Password = hashedPassword

	if requestError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	h.userRepo.Create(h.user)
	w.WriteHeader(http.StatusCreated)
	h.sigResponse.Status = "User Created"
	json.NewEncoder(w).Encode(h.sigResponse)
	return
}

func (h *Handlers)GetUserData(w http.ResponseWriter, request *http.Request){

	uName := request.Context().Value("username") // value from mux context took from ValidateToken middleware

	data := h.userRepo.Find(string(fmt.Sprint(uName)))

	h.user.Id = data.Id
	h.user.Username = data.Username
	h.user.Password = data.Password
	h.user.CreatedAt = data.CreatedAt
	
	data.Password = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
	return
}
