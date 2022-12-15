package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func PaymentHistoryReport (w http.ResponseWriter, r *http.Request){
	var referencePaymentHistory variables.ReferencePaymentHistory
	var response variables.RequestResponse
	err := json.NewDecoder(r.Body).Decode(&referencePaymentHistory)
	if err != nil {
		response.Error = "No se pudo decodificar referencia"
	}
	result := variables.PaymentsReferenceHistory.FindOne(context.TODO(),bson.D{{"reference",referencePaymentHistory.Reference}})
	var JSONresponse []byte
	if result.Err() == mongo.ErrNoDocuments {
		response.Error = "No se encontraron pagos realizados"
		JSONresponse, _ = json.Marshal(response)
	}else {
		err = result.Decode(&referencePaymentHistory)
		if err != nil {
			response.Error = "No se pudo decodificar referencia"
		}else {
			JSONresponse, _ = json.Marshal(referencePaymentHistory.History)
		}
	}
	fmt.Fprintln(w, string(JSONresponse))
}
