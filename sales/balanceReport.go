package sales

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
	"sort"
	"strconv"
	"time"
)

func BalanceReport(w http.ResponseWriter, r *http.Request) {
	var referenceSalesState variables.RefereceSalesState
	var response variables.RequestResponse
	referenceTotalSales := make(map[string]int)
	referencePayedSales := make(map[string]int)
	err := json.NewDecoder(r.Body).Decode(&referenceSalesState)
	if err != nil {
		response.Error = "No se pudo decodificar petición"
	}
	result := variables.ReferencesCollection.FindOne(context.TODO(), bson.D{{"name", referenceSalesState.Name}})
	if result.Err() == mongo.ErrNoDocuments {
		response.Error = "No se encontró resultados"
	} else {
		err = result.Decode(&referenceSalesState)
		if err != nil {
			response.Error = "No se pudo decodificar petición"
		}
		func() {
			var referencePaymentHistory variables.ReferencePaymentHistory
			result = variables.PaymentsReferenceHistory.FindOne(context.TODO(), bson.D{{"reference", referenceSalesState.Name}})
			if result.Err() == mongo.ErrNoDocuments {
				response.Error = "No se existen registros de pagos"
				return
			}
			result.Decode(&referencePaymentHistory)
			sort.Slice(referencePaymentHistory.History, func(i, j int) bool {
				return referencePaymentHistory.History[i].Date > referencePaymentHistory.History[j].Date
			})
			referenceSalesState.Pending = referencePaymentHistory.History[0].Pending
		}()
		cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"reference", referenceSalesState.Name}})
		if err != nil {
			response.Error = "No se existen reservas de proveedor"
		}
		for cursor.Next(context.TODO()) {
			var reserve variables.Reserve
			err = cursor.Decode(&reserve)
			if err != nil {
				log.Fatal(err)
			}
			if reserve.IsPayed == false {
				referenceSalesState.Balance += reserve.Price
			}
			dateToSave := reserve.Date.Time().String()[:7]
			referenceTotalSales[dateToSave] += 1
			if reserve.IsPayed == true {
				referencePayedSales[dateToSave] += 1
			}
		}

		for i, v := range referenceTotalSales {
			var refereceSalesMonth variables.RefereceSalesMonth
			refereceSalesMonth.TotalReserves = v
			if referencePayedSales[i] > 0 {
				refereceSalesMonth.PayedReserves = referencePayedSales[i]
			}

			year, err := strconv.Atoi(i[:4])
			month, err := strconv.Atoi(i[5:7])
			genTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
			if err != nil {
				log.Fatal(err)
			}
			refereceSalesMonth.Date = primitive.NewDateTimeFromTime(genTime)

			referenceSalesState.MonthlyData = append(referenceSalesState.MonthlyData, refereceSalesMonth)
		}
	}
	JSONresponse := []byte{}
	if len(referenceSalesState.MonthlyData) > 0 && referenceSalesState.MonthlyData == nil {
		JSONresponse, _ = json.Marshal(response)
	} else {
		JSONresponse, _ = json.Marshal(referenceSalesState)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
