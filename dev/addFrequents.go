package dev

import (
	"context"
	"fmt"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetFrequents (w http.ResponseWriter, r *http.Request){
	cursor, err := variables.FrequentsCollection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()){
		var frequentOld variables.FrequentsOld
		cursor.Decode(&frequentOld)

		birth := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(),0,0,0,0, time.UTC).AddDate(-frequentOld.Age,0,0)
		fmt.Println(birth)
		frequent := variables.Frequent{
			Id: utils.GenerateID(),
			Name: frequentOld.Name,
			Status: frequentOld.Status,
			Country: fmt.Sprintf("%s%s",frequentOld.Country[:1],strings.ToLower(frequentOld.Country[1:])),
			Passport: frequentOld.Passport,
			Birthday: primitive.NewDateTimeFromTime(birth),
		}
		fmt.Println(frequent)
		res, _ := variables.FrequentsNewCollection.InsertOne(context.TODO(),frequent)
		fmt.Println(res.InsertedID)
	}
}
