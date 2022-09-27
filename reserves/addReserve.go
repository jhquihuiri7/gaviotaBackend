package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"log"
	"net/http"
)

func AddReserves (w http.ResponseWriter, r *http.Request){
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
		reserve := variables.Reserve{}
		err := decoder.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		}
		reserve.Reserve = reserveNumber
		reserve.Id = utils.GenerateID()
		if reserve.Ship == "Gaviota" {
			reservesGaviota = append(reservesGaviota, reserve)
		}else {
			reservesOther = append(reservesOther, reserve)
		}
	}
	if reservesGaviota != nil {
		variables.ReservesGaviotaCollection.InsertMany(context.TODO(),reservesGaviota)
		response.Succes = "Reservas agregadas correctamente"
	}
	if reservesOther != nil {
		variables.ReservesOtherCollection.InsertMany(context.TODO(),reservesOther)
		response.Succes = "Reservas agregadas correctamente"
	}
	JSONresponse, err := json.Marshal(response)
	fmt.Println(string(JSONresponse))
}
