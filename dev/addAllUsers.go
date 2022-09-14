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

func AddAllUsers(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("files/Users.csv")
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
		user := variables.UsersData{}
		user.Id = utils.GenerateID()
		user.Password = "123456"
		user.Name = strings.TrimSpace(v[0])
		user.LastName = strings.TrimSpace(v[1])
		user.Rol = strings.TrimSpace(v[2])
		user.User = strings.ToUpper(strings.TrimSpace(user.Name)[0:1] + strings.TrimSpace(user.LastName))
		_, err := variables.UsersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
	}
}
