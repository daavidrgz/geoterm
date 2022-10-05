package country

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type JsonCountryName struct {
	CommonName   string `json:"common"`
	OfficialName string `json:"official"`
}

type JsonFlag struct {
	Svg string `json:"svg"`
}

type JsonCountry struct {
	Code         string                     `json:"cca3"`
	Name         JsonCountryName            `json:"name"`
	Independent  bool                       `json:"independent"`
	Region       string                     `json:"region"`
	Subregion    string                     `json:"subregion"`
	Population   int                        `json:"population"`
	Capital      []string                   `json:"capital"`
	Flag         JsonFlag                   `json:"flags"`
	Translations map[string]JsonCountryName `json:"translations"`
}

func JsonCountryListToCountry(jsonCountries []JsonCountry) []Country {
	var countries []Country
	for _, jsonCountry := range jsonCountries {
		countries = append(countries, JsonCountryToCountry(jsonCountry))
		fmt.Println("")
	}
	return countries
}

func JsonCountryToCountry(jsonCountry JsonCountry) Country {
	fmt.Println("Processing " + jsonCountry.Name.CommonName)
	return Country{
		Code:         jsonCountry.Code,
		Independent:  jsonCountry.Independent,
		Region:       jsonCountry.Region,
		Subregion:    jsonCountry.Subregion,
		Population:   jsonCountry.Population,
		CommonName:   jsonCountry.Name.CommonName,
		OfficialName: jsonCountry.Name.OfficialName,
		Translations: getTranslations(jsonCountry.Translations),
		Capital:      getCapital(jsonCountry.Capital),
		Flag:         getFlag(jsonCountry.Flag.Svg),
	}
}

func getFlag(url string) []byte {
	fmt.Println("Fetching flag...")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func getCapital(capitals []string) string {
	if len(capitals) == 0 {
		return ""
	}
	return capitals[0]
}

func getTranslations(translationMap map[string]JsonCountryName) []Translation {
	var translationsList []Translation
	for language, translation := range translationMap {
		translationsList = append(translationsList, Translation{
			Language:    strings.ToUpper(language),
			Translation: translation.CommonName,
		})
	}
	return translationsList
}
