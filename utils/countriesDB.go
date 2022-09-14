package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func CuntriesDB(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("files/CountriesData.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = 2
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var countriesDB []variables.Country
	for _, v := range records {
		country := variables.Country{Name: v[0], Code: v[1]}
		countriesDB = append(countriesDB, country)
	}
	JSONCountries, err := json.Marshal(countriesDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONCountries))
}
func DownloadCuntriesDB(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("files/CountriesData.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w.Header().Set("Content-Disposition", "attachment; filename=countries.csv")
	w.Header().Set("Content-Type", r.Header.Get("text/csv"))
	countries, err := ioutil.ReadAll(f)
	if err == io.EOF {
		fmt.Println("Archivo leido completamente")
	}
	w.Write(countries)
}
