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
	databasePath = filepath.Join(homeDir, ".cache/geoterm/countries.db")

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
	create table countries (code text primary key, common text, official text, capital text, region text, flag blob);
	create table translations (code text foreign key references country(code), translation)
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
	insert into countries(code, common, official, region, capital, flag) values(?, ?, ?, ?, ?, ?);
	`

	_, err := database.Exec(sqlStmt, country.Code, country.CommonName, country.OfficialName, country.Continent, country.Capital, country.Flag)
	if err != nil {
		log.Fatal("Error while inserting country")
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

	err := database.QueryRow(sqlStmt, name, name).Scan(&country.Code, &country.CommonName, &country.OfficialName,
		&country.Continent, &country.Capital, &country.Flag)
	if err != nil {
		log.Fatal("Error while getting country by name")
	}

	return country
}

func GetAllCountries() []ct.Country {
	var countries []ct.Country

	sqlStmt := `
	select * from countries;
	`

	rows, err := database.Query(sqlStmt)
	if err != nil {
		log.Fatal("Error while getting all countries")
	}

	for rows.Next() {
		var country ct.Country
		err = rows.Scan(&country.Code, &country.CommonName, &country.OfficialName, &country.Continent,
			&country.Capital, &country.Flag)
		if err != nil {
			log.Fatal("Error while scanning country")
		}
		countries = append(countries, country)
	}

	return countries
}
