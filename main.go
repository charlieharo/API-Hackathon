package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Persona

func GetPersonasEndpoint(w http.ResponseWriter, req *http.Request){
}
func GetPersonaEndpoint(w http.ResponseWriter, req *http.Request){	
}
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
}
func DeletePersonaEndpoint(w http.ResponseWriter, req *http.Request){
}

func main()  {
	router := mux.NewRouter()
	
	//endpoints
	router.HandleFunc("/personas", getPersonasEndpoint).Methods("GET")
	router.HandleFunc("/personas/{id}", getPersonaEndpoint).Methods("GET")
	router.HandleFunc("/personas/{id}", CreatePersonaEndpoint).Methods("POST")
	router.HandleFunc("/personas/{id}", DeletePersonaEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))


}