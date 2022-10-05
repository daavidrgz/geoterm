package country

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	translateEndpoint = "https://api-free.deepl.com/v2/translate"
)

type JsonCountryName struct {
	CommonName   string `json:"common"`
	OfficialName string `json:"official"`
}

type JsonFlag struct {
	Svg string `json:"svg"`
}

type JsonCountry struct {
	Name      JsonCountryName `json:"name"`
	Continent string          `json:"region"`
	Capital   []string        `json:"capital"`
	Code      string          `json:"cca3"`
	Flag      JsonFlag        `json:"flags"`
}

type Country struct {
	Code         string
	CommonName   string
	OfficialName string
	Translations []string
	Continent    string
	Capital      string
	Flag         []byte
}

func fetchFlag(url string) []byte {
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

func JsonCountryListToJson(jsonCountries []JsonCountry) []Country {
	var countries []Country
	for _, jsonCountry := range jsonCountries {
		countries = append(countries, JsonCountryToJson(jsonCountry))
	}
	return countries
}

func JsonCountryToJson(jsonCountry JsonCountry) Country {
	var country Country
	country.Code = jsonCountry.Code
	country.CommonName = jsonCountry.Name.CommonName
	country.OfficialName = jsonCountry.Name.OfficialName
	country.Translations = getTranslations(jsonCountry.Name.CommonName)
	country.Continent = jsonCountry.Continent
	if len(jsonCountry.Capital) == 0 {
		country.Capital = ""
	} else {
		country.Capital = jsonCountry.Capital[0]
	}
	country.Flag = fetchFlag(jsonCountry.Flag.Svg)
	return country
}

func MatchesName(country Country, guess string) bool {
	guess = strings.ToLower(guess)

	for _, translation := range country.Translations {
		if strings.ToLower(translation) == guess {
			return true
		}
	}
	return strings.ToLower(country.CommonName) == guess
}

func MatchesCapital(country Country, capitalName string) bool {
	return country.Capital == strings.ToLower(capitalName)
}

func getTranslations(countryName string) []string {
	targetLangs := []string{"DE", "FR", "ES"}
	var translations []string
	for _, targetLang := range targetLangs {
		translations = append(translations, translate(countryName, "EN", targetLang))
	}
	return translations
}

func translate(text, sourceLang, targetLang string) string {
	req, err := http.NewRequest("POST", translateEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+os.Getenv("DEEPL_API_KEY"))
	req.Header.Set("User-Agent", "geoterm/0.1.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := req.URL.Query()
	q.Add("text", text)
	q.Add("source_lang", sourceLang)
	q.Add("target_lang", targetLang)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var objmap map[string]json.RawMessage
	json.Unmarshal(body, &objmap)

	var translations []json.RawMessage
	json.Unmarshal(objmap["translations"], &translations)

	var result map[string]string
	json.Unmarshal(translations[0], &result)

	return result["text"]
}
