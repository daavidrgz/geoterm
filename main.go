package main

import (
	"geoterm/internal/database"
	"geoterm/pkg/guessflag"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Init()
	// api.FetchCountries()
	guessflag.LaunchGame()
	database.Close()
}
