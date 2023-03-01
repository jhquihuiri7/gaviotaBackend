package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func SaleReport(w http.ResponseWriter, r *http.Request) {
	var filter variables.ReportSalesFilter
	var reservesTemp []variables.Reserve
	var salesSliceReport []variables.ReportSalesData
	var response variables.RequestResponse
	filterMap := make(map[string]int)

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		response.Error = "No es posible decodificar reservas"
	}
	//opts := options.Find().SetProjection(bson.D{{"passenger", 1},{"age", 1}, {"reference", 1}, {"date", 1},{"time", 1},{"price", 1},{"route", 1},{"isPayed", 1},{"ship", 1}})
	var cursor *mongo.Cursor
	if filter.Collection == "Gaviota" {
		cursor, err = variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"date", bson.D{{"$gte", filter.InitDate}, {"$lte", filter.FinalDate}}}, {"isPayed", false}})
	} else {
		cursor, err = variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"date", bson.D{{"$gte", filter.InitDate}, {"$lte", filter.FinalDate}}}})
	}

	if err != nil && cursor.Err() != nil {
		response.Error = "No es posible encontrar reservas"
	}
	values, _ := cursor.Current.Values()
	if len(values) == 0 {
		response.Error = "No es posible encontrar reservas"
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			response.Error = "No reservas encontradas"
		}
		if filter.Collection == "Gaviota" {
			filterMap[reserve.Reference] += reserve.Price
		} else {
			filterMap[reserve.Ship] += reserve.Price
		}
		reservesTemp = append(reservesTemp, reserve)
	}
	for i, _ := range filterMap {
		var saleReportData variables.ReportSalesData
		saleReportData.Reference = i
		//saleReportData.Total = v
		var reserves []variables.Reserve
		for _, re := range reservesTemp {
			if filter.Collection == "Gaviota" {
				if i == re.Reference {
					reserves = append(reserves, re)
					saleReportData.Total += re.Price
				}
			} else {
				if i == re.Ship {
					reserves = append(reserves, re)
					if re.IsPayed {
						saleReportData.Payed += re.Price
					} else {
						saleReportData.Total += re.Price
					}
				}
			}

		}
		saleReportData.Reserves = reserves
		salesSliceReport = append(salesSliceReport, saleReportData)
	}
	var JSONresponse []byte
	if len(salesSliceReport) != 0 {
		JSONresponse, _ = json.Marshal(salesSliceReport)
	} else {
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}

func SaleReportOther(w http.ResponseWriter, r *http.Request) {
	var filter variables.ReportSalesFilter
	var reservesTemp []variables.Reserve
	var salesSliceReport []variables.ReportSalesData
	var response variables.RequestResponse
	filterMap := make(map[string]int)

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		response.Error = "No es posible decodificar reservas"
	}
	opts := options.Find().SetProjection(bson.D{{"passenger", 1}, {"reference", 1}, {"date", 1}, {"time", 1}, {"price", 1}, {"route", 1}, {"isPayed", 1}, {"ship", 1}})
	cursor, err := variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"date", bson.D{{"$gte", filter.InitDate}, {"$lte", filter.FinalDate}}}}, opts)

	if err != nil && cursor.Err() != nil {
		response.Error = "No es posible encontrar reservas"
	}
	values, _ := cursor.Current.Values()
	if len(values) == 0 {
		response.Error = "No es posible encontrar reservas"
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			response.Error = "No reservas encontradas"
		}
		filterMap[reserve.Ship] += reserve.Price
		reservesTemp = append(reservesTemp, reserve)
	}
	for i, v := range filterMap {
		var saleReportData variables.ReportSalesData
		saleReportData.Reference = i
		saleReportData.Total = v
		//result := variables.ReferencesCollection.FindOne(context.TODO(),bson.D{{"name",i}})
		//if result.Err() == mongo.ErrNoDocuments{
		//	response.Error = "No reservas encontradas"
		//}
		//err = result.Decode(&saleReportData)
		//if err != nil {
		//	response.Error = "No es posible encontrar reservas"
		//}
		var reserves []variables.Reserve
		for _, re := range reservesTemp {
			if i == re.Reference {
				reserves = append(reserves, re)
			}
		}
		saleReportData.Reserves = reserves
		salesSliceReport = append(salesSliceReport, saleReportData)
	}
	var JSONresponse []byte
	if len(salesSliceReport) != 0 {
		JSONresponse, _ = json.Marshal(salesSliceReport)
	} else {
		JSONresponse, _ = json.Marshal(response)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
