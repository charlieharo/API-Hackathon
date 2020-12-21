package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func GetPersonnasEndpoint(w http.REsponseWriter, req *http.Request){


}


func ()  {
	router := mux.NewRouter()
	
	//endpoints
	router.HandleFunc("/personas", getPersonasEndpoint).Methods("GET")
}