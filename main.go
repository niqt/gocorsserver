package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gocorsserver/controllers"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/get", controllers.GetEndpoint).Methods("GET")
	router.HandleFunc("/api/v1/create", controllers.CreateEndpoint).Methods("POST")
	

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"accept", "x-authorization", "content-type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	fmt.Printf((http.ListenAndServe(":9090", c.Handler(router))).Error())
}

