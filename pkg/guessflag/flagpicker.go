package guessflag

import (
	"bytes"
	"fmt"
	ct "geoterm/internal/country"
	"geoterm/internal/database"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var countries []ct.Country
var country ct.Country

func Init() {
	rand.Seed(time.Now().UnixNano())
	countries = database.GetAllCountries()
	index := rand.Intn(len(countries))
	country = removeCountry(index)
}

func NextCountry() {
	index := rand.Intn(len(countries))
	country = removeCountry(index)
}

func GetCurrentCountry() ct.Country {
	return country
}

func ShowFlag() {
	fmt.Println(country.Translations)
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
