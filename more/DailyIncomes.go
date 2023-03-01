package more

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetDailyIncome(w http.ResponseWriter, r *http.Request) {
	var request variables.DailyReportRequest
	var response variables.RequestResponse
	var reportData []variables.DailyIncomes
	paymentMethods := []string{
		"Efectivo",
		"Payphone",
		"Transferencia Banco Pichincha",
		"Transferencia Banco Pacífico",
		"Depósito Banco Pichincha",
		"Depósito Banco Pacífico"}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error = "No se pudo decodificar petición"
	}
	gaviotaCursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.M{"date": bson.M{"$eq": request.Date}})
	otherCursor, err := variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"date", request.Date}})
	var reserves []variables.Reserve
	for gaviotaCursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = gaviotaCursor.Decode(&reserve)

		if err != nil {
			response.Error = "No se pudo decodificar reserva"
		}
		reserves = append(reserves, reserve)
	}
	for otherCursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = otherCursor.Decode(&reserve)

		if err != nil {
			response.Error = "No se pudo decodificar reserva"
		}
		reserves = append(reserves, reserve)
	}
	for _, v := range paymentMethods {
		var dailyIncome variables.DailyIncomes
		dailyIncome.Method = v

		for _, val := range reserves {
			if v == val.PaymentMethod {
				dailyIncome.Total += val.Price
				dailyIncome.Reserves = append(dailyIncome.Reserves, val)
			}
		}
		if dailyIncome.Total == 0 {
			continue
		}
		reportData = append(reportData, dailyIncome)
	}
	var JSONresponse []byte
	if len(reportData) > 0 {
		JSONresponse, _ = json.Marshal(reportData)
	} else {
		if response.Error == "" {
			response.Error = "No existen pagos realizados en la fecha seleccionada"
		}
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
