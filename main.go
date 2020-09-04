package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/doduneves/planet-api/config"
	controllers "github.com/doduneves/planet-api/controllers"
	"github.com/gorilla/mux"
)

var config = Config{}

func init() {
	config.Read()
}

func main() {
	config.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/planets", controllers.GetAll).Methods("GET")
	r.HandleFunc("/planets/{id}", controllers.GetByID).Methods("GET")
	r.HandleFunc("/planets", controllers.Create).Methods("POST")
	r.HandleFunc("/planets/{id}", controllers.Delete).Methods("DELETE")

	var port = ":8001"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
