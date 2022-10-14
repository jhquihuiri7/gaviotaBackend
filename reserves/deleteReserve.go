package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func DeleteReserve(w http.ResponseWriter, r *http.Request) {
	var deleteReserves []variables.Reserve
	var response variables.RequestResponse
	decoder := json.NewDecoder(r.Body)
	_, err := decoder.Token()
	if err != nil {
		response.Error = "No fue posible decodificar reservas"
	}
	for decoder.More() {
		var reserve variables.Reserve
		err = decoder.Decode(&reserve)
		if err != nil {
			response.Error = "No fue posible decodificar todas las reservas"
		}else {deleteReserves = append(deleteReserves, reserve)}
	}
	if len(deleteReserves) == 0 {
		response.Error = "No hay reservas por eliminar"
	}
	for _, v:= range deleteReserves {
		var binReserve variables.Reserve
		result := variables.ReservesGaviotaCollection.FindOne(context.TODO(),bson.D{{"_id",v.Id}})
		if result.Err() == mongo.ErrNoDocuments {
			result = variables.ReservesOtherCollection.FindOne(context.TODO(),bson.D{{"_id",v.Id}})
			if result.Err() == mongo.ErrNoDocuments {
				response.Error = "No se encontró reserva"
				continue
			}else {
				err = result.Decode(&binReserve)
				if err != nil {
					response.Error = "No fue posible decodificar reservas"
				}else {
					_, err =variables.BinReservesCollection.InsertOne(context.TODO(),binReserve)
					if err != nil {
						response.Error = "No es posible enviar reserva a papelera"
					}
				}
				count, _ := variables.ReservesOtherCollection.DeleteOne(context.TODO(),bson.D{{"_id",v.Id}})
				if count.DeletedCount == 0 {
					response.Error = "No fue posible eliminar un reserva"
				}else {
					response.Error = ""
				}
			}
		}else {
			err = result.Decode(&binReserve)
			if err != nil {
				response.Error = "No fue posible decodificar reservas"
			}else {
				_, err =variables.BinReservesCollection.InsertOne(context.TODO(),binReserve)
				if err != nil {
					response.Error = "No es posible enviar reserva a papelera"
				}
			}
			count, _ := variables.ReservesGaviotaCollection.DeleteOne(context.TODO(),bson.D{{"_id",v.Id}})
			if count.DeletedCount == 0 {
				response.Error = "No fue posible eliminar un reserva"
			}else {
				response.Error = ""
			}
		}
	}
	if response.Error == "" {
		response.Succes = "Reservas eliminadas correctamente"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
