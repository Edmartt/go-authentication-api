package server

import (
	"fmt"

	"log"
	"net/http"

	"github.com/Edmartt/go-authentication-api/internal/users/transport"

	"github.com/gorilla/mux"
)


type HttpServer struct{
	Router *mux.Router
	handlers transport.Handlers
}

func (server *HttpServer) SetServer(){
	server.Router = mux.NewRouter().StrictSlash(true)
	server.SetRoutes()
}

func (server *HttpServer) StartServer(port string){
	fmt.Println("Server Started in: ", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func (server *HttpServer) SetRoutes(){

	server.Router.HandleFunc("/api/v1/public/login", transport.SetJSONResponse(server.handlers.Login)).Methods("POST")

	server.Router.HandleFunc("/api/v1/public/signup",transport.SetJSONResponse(transport.ValidateRequestBody(server.handlers.Signup))).Methods("POST")

	server.Router.HandleFunc("/api/v1/users/signup",transport.SetJSONResponse(transport.ValidateRequestBody(server.handlers.Signup))).Methods("POST")

}
