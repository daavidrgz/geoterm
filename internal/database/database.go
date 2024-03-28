package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	ct "geoterm/internal/country"
)

var databasePath string
var database *sql.DB

func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error while getting home directory")
	}
	databasePath = filepath.Join(homeDir, ".geoterm/countries.db")

	openDatabase()
}

func Close() {
	database.Close()
}

func existsDatabase() bool {
	_, err := os.Stat(databasePath)
	return !os.IsNotExist(err)
}

func openDatabase() {
	if !existsDatabase() {
		createDatabase()
		return
	}

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal("Error opening database")
	}

	database = db
}

func createDatabase() {
	os.MkdirAll(filepath.Dir(databasePath), os.ModePerm)

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal("Error opening database")
	}

	sqlStmt := `
	create table countries (code text primary key, common text, official text,
	 capital text, region text, subregion text, independent integer, population integer, flag blob);
	create table translations (code text, language text, translation text, primary key (code, language, translation),
	 constraint fk_translation foreign key (code) references countries(code));
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Error while crating database")
	}

	database = db
	fmt.Println("Database created")
}

func InsertCountry(country ct.Country) {
	sqlStmt := `
	insert into countries(code, common, official, region, subregion, independent, population, capital, flag) values(?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	_, err := database.Exec(sqlStmt, country.Code, country.CommonName, country.OfficialName, country.Region,
		country.Subregion, country.Independent, country.Population, country.Capital, country.Flag)
	if err != nil {
		log.Fatal("Error while inserting country")
	}

	for _, translation := range country.Translations {
		InsertTranslation(country.Code, translation)
	}
}

func InsertTranslation(code string, translation ct.Translation) {
	sqlStmt := `
	insert into translations(code, language, translation) values(?, ?, ?);
	`

	_, err := database.Exec(sqlStmt, code, translation.Language, translation.Translation)
	if err != nil {
		log.Fatal("Error while inserting translation")
	}
}

func InsertCountries(countries []ct.Country) {
	for _, country := range countries {
		InsertCountry(country)
	}
}

func GetCountryByName(name string) ct.Country {
	var country ct.Country

	sqlStmt := `
	select * from countries where common = ? or official = ?;
	`

	row := database.QueryRow(sqlStmt, name, name)
	err := row.Scan(&country.Code, &country.CommonName, &country.OfficialName, &country.Capital, &country.Region,
		&country.Subregion, &country.Independent, &country.Population, &country.Flag)
	if err != nil {
		log.Fatal("Error while getting country by name")
	}

	translateSqlStmt := `
	select language, translation from translations where code = ?;
	`

	rows, err := database.Query(translateSqlStmt, country.Code)
	if err != nil {
		log.Fatal("Error while getting translations")
	}

	translations := []ct.Translation{}
	for rows.Next() {
		var translation ct.Translation
		err = rows.Scan(&translation.Language, &translation.Translation)
		if err != nil {
			log.Fatal("Error while getting translation")
		}
		translations = append(translations, translation)
	}

	country.Translations = translations

	return country
}

func GetAllCountries() []ct.Country {
	var countries []ct.Country

	sqlStmt := `
	select common from countries;
	`

	rows, err := database.Query(sqlStmt)
	if err != nil {
		log.Fatal("Error while getting all countries")
	}

	for rows.Next() {
		var countryName string
		err = rows.Scan(&countryName)
		if err != nil {
			log.Fatal("Error while getting country name")
		}
		countries = append(countries, GetCountryByName(countryName))
	}
	return countries
}

func GetAllIndependentCountries() []ct.Country {
	var countries []ct.Country

	sqlStmt := `
	select common from countries where independent = 1;
	`

	rows, err := database.Query(sqlStmt)
	if err != nil {
		log.Fatal("Error while getting all independent countries")
	}

	for rows.Next() {
		var countryName string
		err = rows.Scan(&countryName)
		if err != nil {
			log.Fatal("Error while getting country name")
		}
		countries = append(countries, GetCountryByName(countryName))
	}
	return countries
}
