package country

import (
	"strings"
)

type Translation struct {
	Language    string
	Translation string
}

type Country struct {
	Code         string
	Independent  bool
	Region       string
	Subregion    string
	Population   int
	CommonName   string
	OfficialName string
	Translations []Translation
	Capital      string
	Flag         []byte
}

func MatchesName(country Country, guess string) bool {
	for _, translation := range country.Translations {
		if compareName(translation.Translation, guess) {
			return true
		}
	}
	return compareName(country.CommonName, guess)
}

func MatchesCapital(country Country, guess string) bool {
	return compareName(country.Capital, guess)
}

func compareName(a, b string) bool {
	return normalize(a) == normalize(b)
}

func normalize(text string) string {
	accentsMap := map[string]string{
		"á": "a",
		"é": "e",
		"í": "i",
		"ó": "o",
		"ú": "u",
		"ü": "u",
		"Á": "A",
		"É": "E",
		"Í": "I",
		"Ó": "O",
		"Ú": "U",
		"Ü": "U",
	}

	for accent, letter := range accentsMap {
		text = strings.Replace(text, accent, letter, -1)
	}

	return strings.ToLower(text)
}
