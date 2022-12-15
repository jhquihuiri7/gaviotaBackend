package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
					if oldReserve.IsPayed == false {
						update = append(update, bson.E{"isConfirmed",!oldReserve.IsPayed})

					}
				case "isConfirmed":
					update = append(update, bson.E{typeEdit,!oldReserve.IsConfirmed})
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
					update = append(update, bson.E{typeEdit, !oldReserve.IsPayed})
					if oldReserve.IsPayed == false {
						update = append(update, bson.E{"isConfirmed",!oldReserve.IsPayed})
					}
					case "isConfirmed":
					update = append(update, bson.E{typeEdit, !oldReserve.IsConfirmed})
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
func EditReserveSingle (w http.ResponseWriter, r *http.Request){
	var reserve variables.EditedReserve
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&reserve)
	if err != nil {
		response.Error = "No se puede decodificar reserva"
	}
	var newCollection string
	if reserve.Ship == "Gaviota" {
		newCollection = reserve.Ship
	}else{
		newCollection = "Other"
	}
	if reserve.Collection == newCollection {
		if reserve.Ship == "Gaviota" {
			result, _ := variables.ReservesGaviotaCollection.UpdateOne(context.TODO(),bson.D{{"_id",reserve.Id}},UpdateReserve(reserve))

			if result.ModifiedCount == 0 {
				response.Error = "No se puedo actualizar reserva"
			}else {
				response.Succes = "Reserva modificada correctamente"
			}
		}else {
			result, _ :=variables.ReservesOtherCollection.UpdateOne(context.TODO(),bson.D{{"_id",reserve.Id}},UpdateReserve(reserve))
			if result.ModifiedCount == 0 {
				response.Error = "No se puedo actualizar reserva"
			}else{
				response.Succes = "Reserva modificada correctamente"
			}
		}
	}else {
		//TODO implement structural reserve
		//encuentra inserta y elimina
		if reserve.Collection == "Gaviota" {
			_, err = variables.ReservesOtherCollection.InsertOne(context.TODO(),ConvertReserve(reserve))
			if err != nil {
				response.Error = "No se puedo actualizar reserva"
			}else {
				result, err := variables.ReservesGaviotaCollection.DeleteOne(context.TODO(), bson.D{{"_id",reserve.Id}})
				if err != nil && result.DeletedCount == 0{
					response.Error = "No se puedo actualizar reserva"
				}else{
					response.Succes = "Reserva modificada correctamente"
				}
			}
		}else{
			_, err = variables.ReservesGaviotaCollection.InsertOne(context.TODO(),ConvertReserve(reserve))
			if err != nil {
				response.Error = "No se puedo actualizar reserva"
			}else {
				result, err := variables.ReservesOtherCollection.DeleteOne(context.TODO(), bson.D{{"_id",reserve.Id}})
				if err != nil && result.DeletedCount == 0{
					response.Error = "No se puedo actualizar reserva"
				}else{
					response.Succes = "Reserva modificada correctamente"
				}
			}
		}
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w,string(JSONresponse))
}
func EditReserveExternal (w http.ResponseWriter, r *http.Request){
	var reserve variables.EditedReserve
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&reserve)
	if err != nil {
		response.Error = "No se puede decodificar reserva"
	}
	response = EditExternal(reserve)
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w,string(JSONresponse))
}
func EditReserveExternalBase (w http.ResponseWriter, r *http.Request) {
	var ids variables.Ids
	var reserves []variables.EditedReserve
	var response variables.RequestResponse
	param := r.URL.Query()
	ship := param["ship"][0]
	json.NewDecoder(r.Body).Decode(&ids)
	cursor, err := variables.ReservesExternalCollection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.EditedReserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		}else{
			for _, v := range ids.Ids {
				if reserve.Id == v {
					reserve.Ship = ship
					reserves = append(reserves, reserve)
				}
			}
		}
	}
	for _, v := range reserves {
		response = EditExternal(v)
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w,string(JSONresponse))
}
func EditExternal (reserve variables.EditedReserve) variables.RequestResponse{
	var response variables.RequestResponse
	if reserve.Ship == "Undefined" {
		result, _ := variables.ReservesExternalCollection.UpdateOne(context.TODO(),bson.D{{"_id",reserve.Id}},UpdateReserve(reserve))
		if result.ModifiedCount == 0 {
			response.Error = "No se puedo actualizar reserva"
		}else {
			response.Succes = "Reserva modificada correctamente"
		}
	}else if reserve.Ship == "Gaviota" {
		_, err := variables.ReservesGaviotaCollection.InsertOne(context.TODO(),ConvertReserve(reserve))
		if err != nil {
			response.Error = "No se puedo actualizar reserva"
		}else {
			result, err := variables.ReservesExternalCollection.DeleteOne(context.TODO(), bson.D{{"_id",reserve.Id}})
			if err != nil && result.DeletedCount == 0{
				response.Error = "No se puedo actualizar reserva"
			}else{
				response.Succes = "Reserva modificada correctamente"
			}
		}
	}else {
		_, err := variables.ReservesOtherCollection.InsertOne(context.TODO(),ConvertReserve(reserve))
		if err != nil {
			response.Error = "No se puedo actualizar reserva"
		}else {
			result, err := variables.ReservesExternalCollection.DeleteOne(context.TODO(), bson.D{{"_id",reserve.Id}})
			if err != nil && result.DeletedCount == 0{
				response.Error = "No se puedo actualizar reserva"
			}else{
				response.Succes = "Reserva modificada correctamente"
			}
		}
	}
	return  response
}