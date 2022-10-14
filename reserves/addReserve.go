package reserves

import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
			insertReserve.Id = utils.GenerateID()
			if reserve.Ship == "Gaviota" {
				reservesGaviota = append(reservesGaviota, insertReserve)
			} else {
				reservesOther = append(reservesOther, insertReserve)
			}
		}
	}
	if reservesGaviota != nil {
		variables.ReservesGaviotaCollection.InsertMany(context.TODO(), reservesGaviota)
		response.Succes = "Reservas agregadas correctamente"
		go func() {
			var dialer websocket.Dialer
			conn, _, err := dialer.Dial("wss://websocket-microservice.herokuapp.com/api/wsGaviota", http.Request{}.Header)
			defer conn.Close()
			if err != nil {
				fmt.Println(err)
			}
			err = conn.WriteJSON(reservesGaviota)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	if reservesOther != nil {
		variables.ReservesOtherCollection.InsertMany(context.TODO(), reservesOther)
		response.Succes = "Reservas agregadas correctamente"
	}

	JSONresponse, err := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
