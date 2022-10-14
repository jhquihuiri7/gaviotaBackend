package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func EditReserveBase(w http.ResponseWriter, r *http.Request) {
	typeEdit := r.URL.Query()["typeEdit"][0]
	opts := options.FindOne().SetProjection(bson.D{{typeEdit, 1}})
	var editReserves []variables.Reserve
	var response variables.RequestResponse

	decoder := json.NewDecoder(r.Body)
	_, err := decoder.Token()
	if err != nil {
		response.Error = "No se pudo decodificar reservas"
	}
	for decoder.More() {
		var reserve variables.Reserve
		err = decoder.Decode(&reserve)
		if err != nil {
			response.Error = "No se pudo decodificar reservas"
		}
		editReserves = append(editReserves, reserve)
	}

	for _, v := range editReserves {
		result := variables.ReservesGaviotaCollection.FindOne(context.TODO(),bson.D{{"_id",v.Id}},opts)
		var oldReserve variables.Reserve
		var update bson.D
		if result.Err() == mongo.ErrNoDocuments {
			result = variables.ReservesOtherCollection.FindOne(context.TODO(),bson.D{{"_id",v.Id}},opts)
			if result.Err() == mongo.ErrNoDocuments {
				response.Error = "No se encontró reservas"
			}else {
				err = result.Decode(&oldReserve)
				switch typeEdit {
				case "isPayed":
					update = append(update, bson.E{typeEdit,!oldReserve.IsPayed})
				case "isConfirmed":
					update = append(update, bson.E{typeEdit,!oldReserve.IsConfirmed})
				case "isBlocked":
					update = append(update, bson.E{typeEdit,!oldReserve.IsBlocked})
				}
				updated, err := variables.ReservesOtherCollection.UpdateOne(context.TODO(),bson.D{{"_id",v.Id}},bson.D{{"$set",update}})
				if updated.MatchedCount == 0 && err != nil {
					response.Error = "No se pudo editar reservas"
				}else {
					response.Succes = "Reservas editadas correctamente"
				}
			}
		}else {
			err = result.Decode(&oldReserve)
			if err != nil {
				response.Error = "No se pudo decodificar reservas"
			}else {
				switch typeEdit {
				case "isPayed":
					update = append(update, bson.E{typeEdit,!oldReserve.IsPayed})
				case "isConfirmed":
					update = append(update, bson.E{typeEdit,!oldReserve.IsConfirmed})
				case "isBlocked":
					update = append(update, bson.E{typeEdit,!oldReserve.IsBlocked})
				}
				updated, err := variables.ReservesGaviotaCollection.UpdateOne(context.TODO(),bson.D{{"_id",v.Id}},bson.D{{"$set",update}})
				if updated.MatchedCount == 0 && err != nil {
					response.Error = "No se pudo editar reservas"
				}else {
					response.Succes = "Reservas editadas correctamente"
				}
			}
		}
	}
	if response.Succes != "" {
		response.Error = ""
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
