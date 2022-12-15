package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

func GetReservesDaily(w http.ResponseWriter, r *http.Request) {
	var cursor *mongo.Cursor
	var err error
	var reserveDate variables.Reserve
	var JSONresponse []byte
	err = json.NewDecoder(r.Body).Decode(&reserveDate)
	if err != nil {
		log.Fatal(err)
	}

	book := r.URL.Query()["book"][0]

	if book == "gaviota" {
		cursor, err = variables.ReservesGaviotaCollection.Find(context.TODO(), bson.M{"date": bson.M{
			"$eq": reserveDate.Date,
		}})
	} else {
		cursor, err = variables.ReservesOtherCollection.Find(context.TODO(), bson.M{"date": bson.M{
			"$eq": reserveDate.Date,
		}})
	}
	if err != nil {
		log.Fatal(err)
	}
	var reserves []variables.Reserve
	for cursor.Next(context.TODO()) {
		i := variables.Reserve{}
		cursor.Decode(&i)
		reserves = append(reserves, i)
	}
	if len(reserves) == 0 {
		JSONresponse, err = json.Marshal(variables.RequestResponse{Error: "No reservas encontradas"})
	} else {
		JSONresponse, err = json.Marshal(reserves)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}

func GetReservesExternal(w http.ResponseWriter, r *http.Request) {
	var response variables.RequestResponse
	var reserves []variables.Reserve
	cursor, err := variables.ReservesExternalCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		response.Error = "No hay ventas externas"
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			response.Error = "Problema al cargar ventas externas"
		} else {
			reserves = append(reserves, reserve)
		}
	}

	var JSONresponse []byte
	if len(reserves) > 0 {
		JSONresponse, _ = json.Marshal(reserves)
	} else {
		response.Error = "No hay ventas externas"
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}

func GetReservesRange(w http.ResponseWriter, r *http.Request) {
	cur, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.M{"date": bson.M{
		"$lt": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -2)),
		"$gt": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -10)),
	}})
	if err != nil {
		log.Fatal(err)
	}
	var users []variables.Reserve
	for cur.Next(context.TODO()) {
		i := variables.Reserve{}
		cur.Decode(&i)
		users = append(users, i)
	}
	fmt.Println(users)
}
