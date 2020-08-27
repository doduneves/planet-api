package main

import (
	"fmt"
	"net/http"
)

type Planet struct {
}

func getAllPlanets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieving All Users")
}

func insertPlanet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inserting a new Planet")
}

func updatePlanet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating Planet")
}

func deletePlanet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Removing a Planet")
}

func getPlanetById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Retrieving specific Planet")
}
