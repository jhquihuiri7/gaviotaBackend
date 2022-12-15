package reserves

import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strings"
)

func AddReserves(w http.ResponseWriter, r *http.Request) {
	var reservesGaviota []interface{}
	var reservesOther []interface{}
	var response variables.RequestResponse

	decoder := json.NewDecoder(r.Body)
	_, err := decoder.Token()
	if err != nil {
		log.Fatal(err)
	}
	reserveNumber := utils.GenerateReserve()
	for decoder.More() {
		reserve := variables.MultiplyReserve{}
		err = decoder.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		}
		reserve.ReserveNumber = reserveNumber
		insertReserve := MultiplyReserve(reserve)
		for i := 0 ; i < reserve.Number; i++ {
			if strings.HasSuffix(insertReserve.Passenger,"@"){
				result := variables.FrequentsNewCollection.FindOne(context.TODO(),bson.D{{"passport",insertReserve.Passport}})
				if result.Err() == mongo.ErrNoDocuments {
					if insertReserve.Passenger != "" && insertReserve.Status != "" && insertReserve.Country != "" && insertReserve.Passport != "" && insertReserve.Phone != "" {
						newFrequent := variables.Frequent{
							Id: uuid.NewV4().String(),
							Name: strings.ToUpper(strings.ReplaceAll(insertReserve.Passenger,"@","")),
							Status: insertReserve.Status,
							Country: insertReserve.Country,
							Passport: insertReserve.Passport,
							Birthday: insertReserve.Birthday,
							Age: 0,
							Phone: insertReserve.Phone,
						}
						variables.FrequentsNewCollection.InsertOne(context.TODO(),newFrequent)
						insertReserve.Passenger = newFrequent.Name
					}
				}else {
					insertReserve.Passenger = strings.ToUpper(strings.ReplaceAll(insertReserve.Passenger,"@",""))
				}
			}
			if strings.HasSuffix(insertReserve.Passenger,"#"){
				result := variables.FrequentsNewCollection.FindOne(context.TODO(),bson.D{{"passport",insertReserve.Passport}})
				if result.Err() != mongo.ErrNoDocuments {
					if insertReserve.Passenger != "" && insertReserve.Status != "" && insertReserve.Country != "" && insertReserve.Passport != "" && insertReserve.Phone != "" {
						newFrequent := variables.Frequent{
							Id: uuid.NewV4().String(),
							Name: strings.ToUpper(strings.ReplaceAll(insertReserve.Passenger,"#","")),
							Status: insertReserve.Status,
							Country: strings.ToUpper(insertReserve.Country),
							Passport: insertReserve.Passport,
							Birthday: insertReserve.Birthday,
							Age: 0,
							Phone: insertReserve.Phone,
						}
						update := bson.D{{"$set",bson.D{{"status",newFrequent.Status},{"name",newFrequent.Name},{"country",newFrequent.Country},{"phone",newFrequent.Phone},{"birthday",newFrequent.Birthday}}}}
						variables.FrequentsNewCollection.UpdateOne(context.TODO(),bson.D{{"passport",insertReserve.Passport}},update)
						insertReserve.Passenger = newFrequent.Name
					}else {
						insertReserve.Passenger = strings.ToUpper(strings.ReplaceAll(insertReserve.Passenger,"#",""))
					}
				}
			}
			insertReserve.Id = utils.GenerateID()
			insertReserve.Passenger =  strings.ToUpper(insertReserve.Passenger)
			if reserve.Ship == "Gaviota" {
				reservesGaviota = append(reservesGaviota, insertReserve)
			} else {
				reservesOther = append(reservesOther, insertReserve)
			}
		}
	}
	if reservesGaviota != nil {
		variables.ReservesGaviotaCollection.InsertMany(context.TODO(), reservesGaviota)
		response.Succes = reserveNumber
		//go func() {
		//	var dialer websocket.Dialer
		//	conn, _, err := dialer.Dial("wss://websocket-microservice.herokuapp.com/api/wsGaviota", http.Request{}.Header)
		//	defer conn.Close()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//	err = conn.WriteJSON(reservesGaviota)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}()
	}
	if reservesOther != nil {
		variables.ReservesOtherCollection.InsertMany(context.TODO(), reservesOther)
		response.Succes = reserveNumber
	}

	JSONresponse, err := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}

func AddReservesExternal(w http.ResponseWriter, r *http.Request) {
	var reservesExternal []interface{}
	var response variables.RequestResponse
	decoder := json.NewDecoder(r.Body)
	_, err := decoder.Token()
	if err != nil {
		log.Fatal(err)
	}
	reserveNumber := utils.GenerateReserve()
	for decoder.More() {
		reserve := variables.MultiplyReserve{}
		err = decoder.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		}
		reserve.ReserveNumber = reserveNumber

		insertReserve := MultiplyReserveExternal(reserve)

		for i := 0 ; i < reserve.Number; i++ {
			insertReserve.Id = utils.GenerateID()
			reservesExternal = append(reservesExternal, insertReserve)
		}
	}
	if reservesExternal != nil {
		result, err := variables.ReservesExternalCollection.InsertMany(context.TODO(), reservesExternal)
		if err != nil || len(result.InsertedIDs) == 0 {
			response.Error = "No se pudo agregar reservas"
		}else {
			response.Succes = reserveNumber
		}
	}
	JSONresponse, err := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
