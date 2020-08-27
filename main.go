package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/doduneves/planet-api/config"
	planetcontroller "github.com/doduneves/planet-api/controllers"
	"github.com/gorilla/mux"
)

var config = Config{}

func init() {
	config.Read()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/planets", planetcontroller.GetAll).Methods("GET")
	r.HandleFunc("/planets/{id}", planetcontroller.GetByID).Methods("GET")
	r.HandleFunc("/planets", planetcontroller.Create).Methods("POST")
	r.HandleFunc("/planets/{id}", planetcontroller.Update).Methods("PUT")
	r.HandleFunc("/planets/{id}", planetcontroller.Delete).Methods("DELETE")

	var port = ":8001"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
