package reserves

import (
	"context"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func EditReserveSingleGaviotaLogic(editReserve variables.Reserve, replace bson.D) string {

	updatedCount, err := variables.ReservesGaviotaCollection.UpdateOne(context.TODO(), bson.D{{"_id", editReserve.Id}}, bson.D{{"$set", replace}})
	if err != nil {
		log.Fatal(err)
	}
	if updatedCount.MatchedCount == 0 {
		var updatedReserve variables.Reserve
		err = variables.ReservesOtherCollection.FindOne(context.TODO(), bson.D{{"_id", editReserve.Id}}).Decode(&updatedReserve)
		if err != nil {
			log.Fatal(err)
		}
		if editReserve.Ship == "Gaviota" {
			updatedReserve.Ship = editReserve.Ship
			updatedReserve.Date = editReserve.Date
			updatedReserve.Route = editReserve.Route
			updatedReserve.Time = editReserve.Time
			_, err := variables.ReservesGaviotaCollection.InsertOne(context.TODO(), updatedReserve)
			if err != nil {
				log.Fatal(err)
			} else {
				variables.ReservesOtherCollection.DeleteOne(context.TODO(), bson.D{{"_id", editReserve.Id}})
			}
		}
	}
	return ""
}
