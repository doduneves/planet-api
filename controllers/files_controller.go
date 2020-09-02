package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func readTxtFile(file string) string {

	lastUpdateFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer lastUpdateFile.Close()

	lastUpdateText := "2006-01-02"

	scanner := bufio.NewScanner(lastUpdateFile)
	for scanner.Scan() {
		lastUpdateText = scanner.Text()
	}

	return lastUpdateText
}

func writeTextFile(file string, lastUpdateText string) {

	_ = ioutil.WriteFile(file, []byte(lastUpdateText), 0644)

}

func GetPlanetsAppearances() map[string]int {

	// Verifico a necessidade de atualizar o 'cache' com o numero de filmes que cada planeta aparece
	dateFormat := "2006-01-02"

	today, err := time.Parse(dateFormat, time.Now().Format(dateFormat))
	if err != nil {
		log.Fatalln(err)
	}

	lastUpdateText := readTxtFile("files/last_update.txt")

	lastUpdate, err := time.Parse(dateFormat, lastUpdateText)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(lastUpdateText)

	fileName := "files/planet_appearances.json"

	fmt.Println("Retrieving file " + fileName)

	var result map[string]int

	if lastUpdate.Before(today) && !lastUpdate.Equal(today) {
		// Se precisar atualizar, pego a resposta direção da API

		appearancesMap := make(map[string]int)

		appearancesMap = RequestSwapiToGetAppearances(appearancesMap, "https://swapi.dev/api/planets")

		result = appearancesMap

		newFile, _ := json.MarshalIndent(appearancesMap, "", "\t\t")
		err = ioutil.WriteFile(fileName, newFile, 0644)
		if err != nil {
			log.Fatalln(err)
		}

		err = ioutil.WriteFile("files/last_update.txt", []byte(time.Now().Format(dateFormat)), 0644)
		if err != nil {
			log.Fatalln(err)
		}

	} else {
		// Se não, o retorno vem do arquivo em cache mesmo

		jsonFile, err := os.Open(fileName)

		if err != nil {
			log.Fatal(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &result)

	}

	return result

}
