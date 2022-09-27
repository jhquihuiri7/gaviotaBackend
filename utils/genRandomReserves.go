package utils

import (
	"encoding/csv"
	"fmt"
	"gaviotaBackend/variables"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetRandomReserves(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query()["number"][0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	f, err := os.Open("files/Passengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < n; i++ {
		v := variables.Reserve{}
		fmt.Println(v, i)
		v.RandomReserve(records)
		fmt.Println(v)
	}
}

//func (r variables.Reserve)RandomReserve (n int){}
