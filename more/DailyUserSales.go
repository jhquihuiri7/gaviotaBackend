package more

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetDailyUserSales (w http.ResponseWriter, r *http.Request){
	var request variables.DailyReportRequest
	var response variables.RequestResponse
	var reportData []variables.DailyUserSales
	userReserves := make(map[string][]variables.Reserve)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error = "No se pudo decodificar petición"
	}
	gaviotaCursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(),bson.D{{"date",request.Date}})
	otherCursor, err := variables.ReservesOtherCollection.Find(context.TODO(),bson.D{{"date",request.Date}})
	for gaviotaCursor.Next(context.TODO()){
		var reserve variables.Reserve
		err = gaviotaCursor.Decode(&reserve)
		if err != nil {
			response.Error = "No se pudo decodificar reserva"
		}
		userReserves[reserve.User] = append(userReserves[reserve.User],reserve)
	}

	for otherCursor.Next(context.TODO()){
		var reserve variables.Reserve
		err = otherCursor.Decode(&reserve)
		if err != nil {
			response.Error = "No se pudo decodificar reserva"
		}
		userReserves[reserve.User] = append(userReserves[reserve.User],reserve)
	}

	for i, v := range userReserves {
		dailyUserSales := variables.DailyUserSales{User: i}
		for _, re := range v {
			dailyUserSales.Total += re.Price
			switch re.PaymentMethod {
			case "No realizado":
				dailyUserSales.NotPayed += re.Price
			case "Efectivo":
				dailyUserSales.Cash += re.Price
			case "Tarjeta de crédito o débito":
				dailyUserSales.CreditDebitCard += re.Price
			case "Transferencia Banco Pichincha":
				dailyUserSales.WireTransferPichincha += re.Price
			case "Transferencia Banco Pacífico":
				dailyUserSales.WireTransferPacifico += re.Price
			case "Depósito Banco Pichincha":
				dailyUserSales.DepositPichincha += re.Price
			case "Depósito Banco Pacífico":
				dailyUserSales.DepositPacifico += re.Price
			}
		}
		reportData = append(reportData, dailyUserSales)
	}
	var JSONresponse []byte
	if len(reportData) > 0 {
		JSONresponse, _ = json.Marshal(reportData)
	}else{
		JSONresponse, _ = json.Marshal(response	)
	}
	fmt.Fprintln(w,string(JSONresponse))
}
