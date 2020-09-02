package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	. "github.com/doduneves/planet-api/models"
)

func RequestSwapiToGetAppearances(appearancesMap map[string]int, url string) map[string]int {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for _, resp := range responseObject.Results {
		appearancesMap[resp.Name] = len(resp.Films)
	}

	if responseObject.Next != "" {
		appearancesMap = RequestSwapiToGetAppearances(appearancesMap, responseObject.Next)
	}

	return appearancesMap

}
