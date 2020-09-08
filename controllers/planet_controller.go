package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "github.com/doduneves/planet-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetPlanetAppearancesByName(name string) int {
	var mapAppearances map[string]int
	mapAppearances = GetPlanetsAppearances()

	for key, appNum := range mapAppearances {
		if key == name {
			return appNum
		}
	}

	return 0

}

func GetAll(w http.ResponseWriter, r *http.Request) {

	// Inicialmente faço um tratamento da URL para gerar os links das próximas páginas
	urlRequest := r.Host + r.URL.EscapedPath()

	var planetResponse PlanetResponse

	pageStr := "1"
	if r.URL.Query()["page"] != nil {
		pageStr = r.URL.Query()["page"][0]
	}

	page, erroConv := strconv.Atoi(pageStr)
	if erroConv != nil {
		respondWithError(w, http.StatusInternalServerError, erroConv.Error())
		return
	}

	if page > 1 {
		planetResponse.Links.PreviousPage = urlRequest + "?page=" + strconv.Itoa(page-1)
	}

	planetResponse.Links.NextPage = urlRequest + "?page=" + strconv.Itoa(page+1)
	planetResponse.Links.Actual = urlRequest + "?page=" + strconv.Itoa(page)

	paramNome := ""
	if r.URL.Query()["nome"] != nil {
		paramNome = r.URL.Query()["nome"][0]
		planetResponse.Links.Actual += "&name=" + paramNome
		if planetResponse.Links.NextPage != "" {
			planetResponse.Links.NextPage += "&name=" + paramNome
		}
		if planetResponse.Links.PreviousPage != "" {
			planetResponse.Links.PreviousPage += "&name=" + paramNome
		}
	}

	var planetsList []Planet

	var planet Planet
	planets, err := planet.GetAll(paramNome, page)

	for _, p := range planets {
		p.SetAppearance(GetPlanetAppearancesByName(p.Nome))
		planetsList = append(planetsList, p)
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	planetResponse.Planets = planetsList
	planetResponse.Count = len(planetsList)
	if planetResponse.Count < 10 {
		planetResponse.Links.NextPage = ""
	}

	respondWithJson(w, http.StatusOK, planetResponse)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var planet Planet
	planet, err := planet.GetByID(params["id"])
	planet.SetAppearance(GetPlanetAppearancesByName(planet.Nome))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Planet ID")
		return
	}
	respondWithJson(w, http.StatusOK, planet)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	planet.ID = bson.NewObjectId()
	if err := planet.Create(planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	planet.SetAppearance(GetPlanetAppearancesByName(planet.Nome))
	respondWithJson(w, http.StatusCreated, planet)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet Planet

	if err := planet.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
