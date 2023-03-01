package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"time"
)

func GetRecentAddedReserves(w http.ResponseWriter, r *http.Request) {
	var reserves []variables.Reserve
	var response variables.RequestResponse
	var JSONresponse []byte
	user := strings.ToUpper(r.URL.Query()["user"][0])
	now := time.Now()
	starDate := primitive.NewDateTimeFromTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local))
	endDate := primitive.NewDateTimeFromTime(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local))
	fmt.Println(starDate.Time())
	fmt.Println(endDate.Time())
	opts := options.Find().SetSort(bson.D{{"registeredDate", -1}}).SetLimit(20)
	cursor1, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"user", user}}, opts)
	//cursor1, err := variables.ReservesGaviotaCollection.Find(context.TODO(),bson.D{{"user",user},{"registeredDate",bson.D{{"$gte",starDate},{"$lt",endDate}}}})
	if err != nil {
		response.Error = "Error al procesar petición"
	}
	cursor2, err := variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"user", user}}, opts)
	if err != nil {
		response.Error = "Error al procesar petición"
	}
	for cursor1.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor1.Decode(&reserve)
		if err != nil {
			response.Error = "No se pudo decodificar reservas"
		}
		reserves = append(reserves, reserve)
	}
	for cursor2.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor2.Decode(&reserve)
		if err != nil {
			response.Error = "No se pudo decodificar reservas"
		}
		reserves = append(reserves, reserve)
	}
	if len(reserves) > 0 {
		JSONresponse, _ = json.Marshal(reserves)
	} else {
		if response.Error == "" {
			response.Error = "No se encontraron reservas"
		}
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
