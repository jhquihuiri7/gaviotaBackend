package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

type data struct {
	Id         string `json:"id"`
	ClientTxId string `json:"clientTxId"`
}

func PaymentResponse(w http.ResponseWriter, r *http.Request) {
	var response variables.RequestResponse
	clientTransactionId := r.URL.Query().Get("clientTransactionId")
	payload := data{Id: r.URL.Query().Get("id"), ClientTxId: clientTransactionId}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	url := "https://pay.payphonetodoesposible.com/api/button/V2/Confirm"
	method := "POST"
	client := &http.Client{}
	payloadData := strings.NewReader(string(payloadJSON))

	req, err := http.NewRequest(method, url, payloadData)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", Authorization)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var paymentInfo variables.PaymentInfo
	err = json.NewDecoder(res.Body).Decode(&paymentInfo)
	if err != nil {
		log.Fatal(err)
	}
	paymentInfo.Id = uuid.NewV4().String()
	_, err = variables.PaymentsSystemHistory.InsertOne(context.TODO(), paymentInfo)
	if err != nil {
		response.Error = "No se pudo registrar pago"
	} else {
		response.Succes = "Pago registrado correctamnte"
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}

func GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	var response variables.RequestResponse
	cursor, err := variables.PaymentsSystemHistory.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var paymentsHistory []variables.PaymentInfo
	for cursor.Next(context.TODO()) {
		var paymentInfo variables.PaymentInfo
		err = cursor.Decode(&paymentInfo)
		if err != nil {
			response.Error = "Error al decodificar respuesta"
		}
		paymentsHistory = append(paymentsHistory, paymentInfo)
	}
	var JSONresponse []byte
	if len(paymentsHistory) > 0 {
		JSONresponse, _ = json.Marshal(paymentsHistory)
	} else {
		response.Error = "No se encuentran pagos realizados"
		JSONresponse, _ = json.Marshal(response)
	}

	fmt.Fprintln(w, string(JSONresponse))
}
