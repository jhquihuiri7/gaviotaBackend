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

func DeleteReserve (w http.ResponseWriter, r *http.Request){
	var deleteReserve variables.Reserve
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&deleteReserve)
	if err != nil {
		log.Fatal(err)
	}
	if deleteReserve.Id != ""{
		var result *mongo.SingleResult
		result = variables.ReservesGaviotaCollection.FindOne(context.TODO(),bson.D{{"_id",deleteReserve.Id}})
		if result.Err() == mongo.ErrNoDocuments {
			result = variables.ReservesOtherCollection.FindOne(context.TODO(),bson.D{{"_id",deleteReserve.Id}})
			if result.Err() == mongo.ErrNoDocuments {
				response.Error = "No se encontró registro"
			}else {
				_, err = variables.ReservesOtherCollection.DeleteOne(context.TODO(),bson.D{{"_id",deleteReserve.Id}})
				if err != nil { response.Error = "No se pudo eliminar registro" }
			}
		}else {
			_, err = variables.ReservesGaviotaCollection.DeleteOne(context.TODO(),bson.D{{"_id",deleteReserve.Id}})
			if err != nil { response.Error = "No se pudo eliminar registro" }
		}

		if response.Error == "" {
			err = result.Decode(&deleteReserve)
			if err != nil {
				response.Error = "No se pudo agregar a papelera"
			}else {
				_, err = variables.BinReservesCollection.InsertOne(context.TODO(),deleteReserve)
				if err != nil { response.Error = "No se pudo agregar a papelera" }else {response.Succes = "Reserva eliminada con éxito"}
			}
		}
	}else {
		var cursor *mongo.Cursor
		var reserves []interface{}
		cursor, _ = variables.ReservesGaviotaCollection.Find(context.TODO(),bson.D{{"reserve",deleteReserve.Reserve}})
		fmt.Println(cursor.RemainingBatchLength())
		if cursor.RemainingBatchLength() == 0 {
			response.Error = "No se encontró registros"
		}else {
			_, err = variables.ReservesGaviotaCollection.DeleteMany(context.TODO(),bson.D{{"reserve",deleteReserve.Reserve}})
			if err != nil { response.Error = "No se pudo eliminar registro"}
		}
		for cursor.Next(context.TODO()) {
			i := variables.Reserve{}
			err = cursor.Decode(&i)
			if err != nil { response.Error = "No se pudo agregar a papelera"}
			reserves = append(reserves, i)
		}

		cursor, err = variables.ReservesOtherCollection.Find(context.TODO(),bson.D{{"reserve",deleteReserve.Reserve}})
		if cursor.RemainingBatchLength() == 0 {
			if response.Error == "" {response.Error = "No se encontró registros"}
		}else {
			_, err = variables.ReservesOtherCollection.DeleteMany(context.TODO(),bson.D{{"reserve",deleteReserve.Reserve}})
			if err != nil { response.Error = "No se pudo eliminar registro"}
		}
		for cursor.Next(context.TODO()) {
			i := variables.Reserve{}
			err = cursor.Decode(&i)
			if err != nil { response.Error = "No se pudo agregar a papelera"}
			reserves = append(reserves, i)
		}
		if len(reserves) > 0 {
			_, err = variables.BinReservesCollection.InsertMany(context.TODO(),reserves)
			if err != nil { response.Error = "No se pudo agregar a papelera"}else {response.Succes = "Reservas eliminada con éxito"}
		}
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
