package api

import (
	"io"
	"log"
	"net/http"

	db "geoterm/internal/database"
	js "geoterm/internal/json"
)

const (
	countriesEndpoint = "https://restcountries.com/v3.1"
)

func FetchCountries() {
	res, err := http.Get(countriesEndpoint + "/all")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	countries := js.JsonToCountries(body)
	db.InsertCountries(countries)
}
