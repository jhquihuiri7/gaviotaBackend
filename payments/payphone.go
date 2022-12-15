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
	req.Header.Add("Authorization", "Bearer 98IVdcZW4yAub2AjI_jXsVh-GoQ-oZvtQHIo50CvrBKfwS2nw7WxH9MK63YiDCTZJlSRa6e4bNgTv5h3IM1tvIR1letrEme4qjsXgB8Qx8o8OJF4VWwXSvcARHBtcjCnYKY_osIrQzDM4u6QCnls3eseL6636xmj2ZiqS6NLIRGN15_3N5WgYQDZZ3XdjVS3CuMeAatDm1_mvfGwr3ys804lBTVJggg8aWhA_DmcICPUDeFsmV3xhuuKece9UUnJ5bfg31ibXaj6cg9hOi-t8qoaLz6WqsINSRCOwDO8b60RiivxGjEHtENqReyW_0-gaEx5fg")
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

//sqlStatement := `CREATE TABLE HistoricalPayments (
//    ID INTEGER PRIMARY KEY AUTOINCREMENT,
//    email varchar(255),
//    cardType varchar(255),
//    bin varchar(255),
//    lastDigits varchar(255),
//    deferredCode varchar(255),
//    deferred INTEGER,
//    cardBrandCode varchar(255),
//    cardBrand varchar(255),
//    amount INTEGER,
//    clientTransactionId varchar(255),
//    phoneNumber varchar(255),
//    statusCode INTEGER,
//    transactionStatus varchar(255),
//    authorizationCode varchar(255),
//    messageCode INTEGER,
//    transactionId INTEGER,
//    document varchar(255),
//    currency varchar(255),
//    optionalParameter1 varchar(255),
//    optionalParameter2 varchar(255),
//    optionalParameter3 varchar(255),
//    optionalParameter4 varchar(255),
//    storeName varchar(255),
//    date varchar(255),
//    regionIso varchar(255),
//    transactionType varchar(255),
//    reference varchar(255)
//);`
//_, err = db.Exec(sqlStatement)
//if err != nil {
//log.Printf("%q: %s\n", err, sqlStatement)
//return
//}
