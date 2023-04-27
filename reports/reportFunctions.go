package reports

import (
	"context"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func GetDailyReportNODELETED(DailyRequest variables.DailyReportRequest) []variables.Reserve {
	var Reserves []variables.Reserve
	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"time", DailyRequest.Time}, {"date", DailyRequest.Date}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			Reserves = append(Reserves, reserve)
		}
	}
	return Reserves
}
func FormatDate(date primitive.DateTime) string {
	day := date.Time().UTC().Day()
	month := ""
	year := date.Time().Year()

	switch date.Time().Month() {
	case 1:
		month = "ENERO"
	case 2:
		month = "FEBRERO"
	case 3:
		month = "MARZO"
	case 4:
		month = "ABRIL"
	case 5:
		month = "MAYO"
	case 6:
		month = "JUNIO"
	case 7:
		month = "JULIO"
	case 8:
		month = "AGOSTO"
	case 9:
		month = "SEPTIEMBRE"
	case 10:
		month = "OCTUBRE"
	case 11:
		month = "NOVIEMBRE"
	case 12:
		month = "DICIEMBRE"
	}
	return fmt.Sprintf("FECHA: %d DE %s DE %d", day, month, year)
}
func FormatTime(time string) string {
	if time == "AM" {
		return "7:00 AM"
	} else {
		return "3:00 PM"
	}
}
func FormatRoute(time string) string {
	if time == "AM" {
		return "RUTA: SAN CRISTOBAL - SANTA CRUZ"
	} else {
		return "RUTA: SANTA CRUZ - SAN CRISTOBAL"
	}
}
func FormatDailyName(time string, date primitive.DateTime) string {
	return fmt.Sprintf("%d%d%d %s.pdf", date.Time().Year(), date.Time().Month(), date.Time().Day(), time)
}
