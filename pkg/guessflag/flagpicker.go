package guessflag

import (
	"bytes"
	ct "geoterm/internal/country"
	"geoterm/internal/database"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var countries []ct.Country
var country ct.Country

func InitFlagSystem() int {
	rand.Seed(time.Now().UnixNano())
	countries = database.GetAllIndependentCountries()
	index := rand.Intn(len(countries))
	country = removeCountry(index)

	return len(countries)
}

func NextCountry() {
	index := rand.Intn(len(countries))
	country = removeCountry(index)
}

func GetCurrentCountry() ct.Country {
	return country
}

func ShowFlag() {
	cmd := exec.Command("chafa", "-C", "on", "-s", "60x20")
	cmd.Stdin = bytes.NewReader(country.Flag)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func removeCountry(index int) ct.Country {
	removedCoutry := countries[index]
	countries[index] = countries[len(countries)-1]
	countries = countries[:len(countries)-1]
	return removedCoutry
}
