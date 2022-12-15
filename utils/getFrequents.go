package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetFrequents (w http.ResponseWriter, r *http.Request){
	 var response variables.RequestResponse
	 cursor, err := variables.FrequentsNewCollection.Find(context.TODO(),bson.D{})
	if err != nil {
		response.Error = "No se encontraron pasajeros frecuentes"
	}
	var frequents []variables.Frequent
	for cursor.Next(context.TODO()){
		var frequent variables.Frequent
		err = cursor.Decode(&frequent)
		if err != nil {
			response.Error = "No se pudo decodificar el registro"
		}else {
			frequent.Age = time.Now().AddDate(-frequent.Birthday.Time().Year(),-int(frequent.Birthday.Time().Month()),-frequent.Birthday.Time().Day()).Year()
			frequents = append(frequents, frequent)
		}
	}
	var JSONresponse []byte
	if len(frequents) > 0 {
		JSONresponse, _ = json.Marshal(frequents)
	}else {
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w,string(JSONresponse))
}
