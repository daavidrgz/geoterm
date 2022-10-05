package json

import (
	"encoding/json"
	ct "geoterm/internal/country"
)

func JsonToCountries(jsonBytes []byte) []ct.Country {
	var countries []ct.JsonCountry
	json.Unmarshal(jsonBytes, &countries)
	return ct.JsonCountryListToCountry(countries)
}
