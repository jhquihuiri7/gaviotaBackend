package reserves

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func EditReserve (w http.ResponseWriter, r *http.Request){
	var editReserve variables.Reserve
	oldRoute := r.URL.Query()["oldRoute"][0]

	err := json.NewDecoder(r.Body).Decode(&editReserve)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(editReserve.Date)
	replace := bson.D{}
	replace = append(replace, bson.E{"date", editReserve.Date})
	if editReserve.Time != "" {
		replace = append(replace, bson.E{"time", editReserve.Time})
	}
	if editReserve.Ship != "" {
		replace = append(replace, bson.E{"ship", editReserve.Ship})
	}
	if editReserve.Route != "" {
		replace = append(replace, bson.E{"route", editReserve.Route})
	}
	fmt.Println(editReserve)
	if editReserve.Id != "" {
		//TODO implement  one edit


	}else {
		//TODO implemente all edit
		if ((editReserve.Time == "AM" && editReserve.Route == "SC-SX") || (editReserve.Time == "PM" && editReserve.Route == "SX-SC"))&& editReserve.Ship == "Gaviota"{
			_, err = variables.ReservesGaviotaCollection.UpdateMany(context.TODO(),bson.D{{"reserve",editReserve.Reserve},{"route",oldRoute}},bson.D{{"$set", replace}})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(editReserve.Time,":",editReserve.Route, ":", editReserve.Ship, "Status: OK")
		}else {
			//TODO get , delete , insert
		}
	}
}