package dev

import (
	"context"
	"encoding/csv"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"log"
	"net/http"
	"os"
	"strings"
)

func AddAllReferences(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("files/References.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range records {
		if i == 0 {
			continue
		}
		reference := variables.ReferencesData{}
		reference.Id = utils.GenerateID()
		reference.Name = strings.TrimSpace(v[0])
		_, err := variables.ReferencesCollection.InsertOne(context.TODO(), reference)
		if err != nil {
			log.Fatal(err)
		}
	}
}
