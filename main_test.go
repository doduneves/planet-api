package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/doduneves/planet-api/controllers"
	models "github.com/doduneves/planet-api/models"
	"gopkg.in/mgo.v2/bson"
)

var planetID bson.ObjectId

func TestGetAllPlanets(t *testing.T) {
	req, err := http.NewRequest("GET", "/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong with status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"links":{"actual":"/planets?page=1","next_page":"","previous_page":""},"count":1,"planets":[{"id":"5f5245299344950b901bfbbd","nome":"Teste","clima":"Árido","terreno":"Deserto","aparicoes":0}]}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

func TestCreatePlanet(t *testing.T) {

	var jsonStr = []byte(`{"nome":"Planeta Teste","clima":"Úmido","terreno":"Plano"}`)

	req, err := http.NewRequest("POST", "/planets", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var returnMap models.Planet
	errJSON := json.NewDecoder(rr.Body).Decode(&returnMap)
	if errJSON != nil {
		t.Fatal(errJSON)
	}

	if rr.Body.String() != "" {
		if returnMap.Nome != "Planeta Teste" {
			t.Errorf("Some error inserting key 'nome'")

		}
		if returnMap.Clima != "Úmido" {
			t.Errorf("Some error inserting key 'nome'")

		}
		if returnMap.Terreno != "Plano" {
			t.Errorf("Some error inserting key 'terreno'")
		}
	}

	planetID = returnMap.ID

}

func TestGetPlanetById(t *testing.T) {
	fmt.Println("/planets/" + planetID.Hex())

	req, err := http.NewRequest("GET", fmt.Sprintf("planets/%s", planetID.Hex()), nil)

	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
