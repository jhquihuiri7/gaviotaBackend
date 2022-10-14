package reserves

import (
	"context"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func EditReserveSingleOtherLogic(editReserve variables.Reserve, replace bson.D) string {
	updatedCount, err := variables.ReservesOtherCollection.UpdateOne(context.TODO(), bson.D{{"_id", editReserve.Id}}, bson.D{{"$set", replace}})
	if err != nil {
		log.Fatal(err)
	}
	if updatedCount.MatchedCount == 0 {
		var updatedReserve variables.Reserve
		err = variables.ReservesGaviotaCollection.FindOne(context.TODO(), bson.D{{"_id", editReserve.Id}}).Decode(&updatedReserve)
		if err != nil {
			log.Fatal(err)
		}
		if editReserve.Ship != "Gaviota" {
			_, err := variables.ReservesOtherCollection.InsertOne(context.TODO(), updatedReserve)
			if err != nil {
				log.Fatal(err)
			} else {
				variables.ReservesGaviotaCollection.DeleteOne(context.TODO(), bson.D{{"_id", editReserve.Id}})
			}
		}
	}
	return ""
}
