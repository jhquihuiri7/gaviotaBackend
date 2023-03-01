package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"sort"
	"time"
)

func RegisterAdvance(w http.ResponseWriter, r *http.Request) {
	var advanceDataRequest variables.AdvanceDataRequest
	var advanceData variables.AdvanceData
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&advanceDataRequest)
	if err != nil {
		log.Fatal(err)
		response.Error = "No se puede decodificar pago"
	} else {
		advanceData.PaymentMethod = advanceDataRequest.PaymentMethod
		advanceData.Advance = advanceDataRequest.Advance
		advanceData.User = advanceDataRequest.User
		advanceData.Date = primitive.NewDateTimeFromTime(time.Now().Local())

		var referencePaymentHistory variables.ReferencePaymentHistory
		result := variables.PaymentsReferenceHistory.FindOne(context.TODO(), bson.D{{"reference", advanceDataRequest.ReferenceName}})
		if result.Err() == mongo.ErrNoDocuments && result.Err() != nil {
			advanceData.Total = advanceData.Advance
		} else {
			result.Decode(&referencePaymentHistory)
			if len(referencePaymentHistory.History) > 1 {
				sort.Slice(referencePaymentHistory.History, func(i, j int) bool {
					return referencePaymentHistory.History[i].Date > referencePaymentHistory.History[j].Date
				})
			}
			advanceData.Total = advanceData.Advance + referencePaymentHistory.History[0].Pending
		}

		advanceData.Pending = advanceData.Total
	}

	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"reference", advanceDataRequest.ReferenceName}, {"isPayed", false}})
	if err != nil {
		log.Fatal(err)
		response.Error = "No se encontraron reservas para el pago"
	}
	var reserves []variables.Reserve
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err := cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
			response.Error = "No se puede decodificar reservas para pago"
		}
		reserves = append(reserves, reserve)
	}
	sort.Slice(reserves, func(i, j int) bool {
		return reserves[i].Date < reserves[j].Date
	})
	pursue := ""
	for _, v := range reserves {
		if pursue == "" {
			if float64(advanceData.Pending)/float64(v.Price) >= 1 {
				result, err := variables.ReservesGaviotaCollection.UpdateOne(context.TODO(), bson.D{{"_id", v.Id}}, bson.D{{"$set", bson.D{{"isPayed", true}}}})
				if err != nil {
					log.Fatal(err)
				}
				advanceData.PayedIds = append(advanceData.PayedIds, v.Id)
				fmt.Println(result.ModifiedCount)
				advanceData.Pending -= v.Price
			} else {
				advanceData.Pending = advanceData.Pending % v.Price
				advanceData.Balance += v.Price
				pursue = "balance"
			}
		} else {
			advanceData.Balance += v.Price
		}
	}
	referencePaymentHistory := variables.ReferencePaymentHistory{
		Id:        uuid.NewV4().String(),
		Reference: advanceDataRequest.ReferenceName,
		History:   []variables.AdvanceData{advanceData},
	}
	opts := options.Update().SetUpsert(true)
	result, err := variables.PaymentsReferenceHistory.UpdateOne(context.TODO(), bson.D{{"reference", referencePaymentHistory.Reference}}, bson.D{{"$push", bson.D{{"history", advanceData}}}}, opts)
	if result.ModifiedCount > 0 {
		response.Error = ""
		response.Succes = "Pago registrado correctamente"
	}
	var JSONresponse []byte
	JSONresponse, _ = json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
