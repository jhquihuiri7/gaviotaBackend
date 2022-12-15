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

func GetBin(w http.ResponseWriter, r *http.Request) {
	var response variables.RequestResponse
	var reserves []variables.Reserve
	cursor, err := variables.BinReservesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		response.Error = "Papelera Vacía"
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			response.Error = "Problema al actualizar papelera"
		} else {
			reserves = append(reserves, reserve)
		}
	}

	var JSONresponse []byte
	if len(reserves) > 0 {
		JSONresponse, _ = json.Marshal(reserves)
	} else {
		response.Error = "Papelera Vacía"
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
func RestoreBin(w http.ResponseWriter, r *http.Request) {
	var reserve variables.Reserve
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&reserve)
	if err != nil {
		response.Error = "Error al decodificar"
	}
	err = variables.BinReservesCollection.FindOne(context.TODO(), bson.D{{"_id", reserve.Id}}).Decode(&reserve)
	if err == mongo.ErrNoDocuments {
		response.Error = "No se pudo encontrar reserva"
	} else {
		if reserve.Ship == "Gaviota" {
			_, err = variables.ReservesGaviotaCollection.InsertOne(context.TODO(), reserve)
			if err != nil {
				response.Error = "No se pudo restaurar reserva"
			} else {
				_, err = variables.BinReservesCollection.DeleteOne(context.TODO(), bson.D{{"_id", reserve.Id}})
				if err != nil {
					response.Error = "No se pudo eliminar de papelera"
				} else {
					response.Succes = "Reserva restaurada correctamente"
				}
			}
		} else if reserve.Ship == "Undefined"{
			_, err = variables.ReservesExternalCollection.InsertOne(context.TODO(), reserve)
			if err != nil {
				response.Error = "No se pudo restaurar reserva"
			} else {
				_, err = variables.BinReservesCollection.DeleteOne(context.TODO(), bson.D{{"_id", reserve.Id}})
				if err != nil {
					response.Error = "No se pudo eliminar de papelera"
				} else {
					response.Succes = "Reserva restaurada correctamente"
				}
			}
		} else {
			_, err = variables.ReservesOtherCollection.InsertOne(context.TODO(), reserve)
			if err != nil {
				response.Error = "No se pudo restaurar reserva"
			} else {
				_, err = variables.BinReservesCollection.DeleteOne(context.TODO(), bson.D{{"_id", reserve.Id}})
				if err != nil {
					response.Error = "No se pudo eliminar de papelera"
				} else {
					response.Succes = "Reserva restaurada correctamente"
				}
			}
		}
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
func CleanBin(w http.ResponseWriter, r *http.Request) {
	var response variables.RequestResponse
	_, err := variables.BinReservesCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		response.Error = "Error al limpiar papelera"
	} else {
		response.Succes = "Papelera vaciada completamente"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
