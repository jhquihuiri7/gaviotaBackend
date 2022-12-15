package reserves

import (
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

func MultiplyReserve (multiplyReserve variables.MultiplyReserve)variables.Reserve{
	age := time.Now().AddDate(-multiplyReserve.Birthday.Time().Year(), int(-multiplyReserve.Birthday.Time().Month()),-multiplyReserve.Birthday.Time().Day())
	ageInt := age.Year()
	if age.Year() < 0 {
		ageInt = 0
	}
	return variables.Reserve{
		Id:            multiplyReserve.Id,
		ReserveNumber: multiplyReserve.ReserveNumber,
		Passenger:     multiplyReserve.Passenger,
		Reference:     multiplyReserve.Reference,
		User:          strings.ToUpper(multiplyReserve.User),
		Age:           ageInt,
		Date:          multiplyReserve.Date,
		Time:          multiplyReserve.Time,
		Passport:      strings.ToUpper(multiplyReserve.Passport),
		Country:       multiplyReserve.Country,
		Price:         multiplyReserve.Price,
		Ship:          multiplyReserve.Ship,
		Route:         multiplyReserve.Route,
		IsConfirmed:   multiplyReserve.IsConfirmed,
		IsPayed:multiplyReserve.IsPayed,
		Birthday: multiplyReserve.Birthday,
		Status: multiplyReserve.Status,
		Comment: strings.ToUpper(multiplyReserve.Comment),
		Notes: strings.ToUpper(multiplyReserve.Notes),
		Phone: multiplyReserve.Phone,
		PaymentMethod: multiplyReserve.PaymentMethod,
		PaymentDate: multiplyReserve.PaymentDate,
		RegisteredDate: primitive.NewDateTimeFromTime(time.Now()),
	}
}
func MultiplyReserveExternal (multiplyReserve variables.MultiplyReserve)variables.Reserve{
	age := time.Now().AddDate(-multiplyReserve.Birthday.Time().Year(), int(-multiplyReserve.Birthday.Time().Month()),-multiplyReserve.Birthday.Time().Day())
	price := 0
	ageInt := age.Year()
	if age.Year() < 0 {
		ageInt = 0
	}
	if ageInt >= 0 && ageInt < 2 {
		price = 0
	}else if ageInt >=2 && ageInt < 65 {
		price = 30
	}else {
		price = 20
	}
	return variables.Reserve{
		Id:            multiplyReserve.Id,
		ReserveNumber: multiplyReserve.ReserveNumber,
		Passenger:     multiplyReserve.Passenger,
		Reference:     "External",
		User:          strings.ToUpper("External"),
		Age:           ageInt,
		Date:          multiplyReserve.Date,
		Time:          multiplyReserve.Time,
		Passport:      strings.ToUpper(multiplyReserve.Passport),
		Country:       multiplyReserve.Country,
		Price:         price,
		Ship:          "Undefined",
		Route:         multiplyReserve.Route,
		IsConfirmed:   true,
		IsPayed:false,
		Birthday: multiplyReserve.Birthday,
		Status: multiplyReserve.Status,
		Comment: strings.ToUpper(multiplyReserve.Comment),
		Notes: strings.ToUpper("Pending"),
		Phone: multiplyReserve.Phone,
		PaymentMethod: "No realizado",
		PaymentDate: multiplyReserve.PaymentDate,
		RegisteredDate: primitive.NewDateTimeFromTime(time.Now()),
	}
}
func UpdateReserve (editedReserve variables.EditedReserve)bson.D{
	age := time.Now().AddDate(-editedReserve.Birthday.Time().Year(), int(-editedReserve.Birthday.Time().Month()),-editedReserve.Birthday.Time().Day())
	ageInt := age.Year()
	if age.Year() < 0 {
		ageInt = 0
	}
	return bson.D{{"$set",
		bson.D{
			{"passenger",strings.ToUpper(editedReserve.Passenger)},
			{"reference",editedReserve.Reference},
			{"date",editedReserve.Date},
			{"time",editedReserve.Time},
			{"passport",strings.ToUpper(editedReserve.Passport)},
			{"country",editedReserve.Country},
			{"price",editedReserve.Price},
			{"route",editedReserve.Route},
			{"isConfirmed",editedReserve.IsConfirmed},
			{"isPayed",editedReserve.IsPayed},
			{"birthday",editedReserve.Birthday},
			{"age",ageInt},
			{"status",editedReserve.Status},
			{"comment",strings.ToUpper(editedReserve.Comment)},
			{"notes",strings.ToUpper(editedReserve.Notes)},
			{"ship",editedReserve.Ship},
			{"paymentMethod",editedReserve.PaymentMethod},
			{"paymentDate",editedReserve.PaymentDate},
			{"phone",editedReserve.Phone},
		}},
	}
}
func ConvertReserve (multiplyReserve variables.EditedReserve)variables.Reserve{
	age := time.Now().AddDate(-multiplyReserve.Birthday.Time().Year(), int(-multiplyReserve.Birthday.Time().Month()),-multiplyReserve.Birthday.Time().Day())

	ageInt := age.Year()
	if age.Year() < 0 {
		ageInt = 0
	}
	return variables.Reserve{
		Id:            multiplyReserve.Id,
		ReserveNumber: multiplyReserve.ReserveNumber,
		Passenger:     strings.ToUpper(multiplyReserve.Passenger),
		Reference:     multiplyReserve.Reference,
		User:          strings.ToUpper(multiplyReserve.User),
		Age:           ageInt,
		Date:          multiplyReserve.Date,
		Time:          multiplyReserve.Time,
		Passport:      strings.ToUpper(multiplyReserve.Passport),
		Country:       multiplyReserve.Country,
		Price:         multiplyReserve.Price,
		Ship:          multiplyReserve.Ship,
		Route:         multiplyReserve.Route,
		IsConfirmed:   multiplyReserve.IsConfirmed,
		IsPayed:multiplyReserve.IsPayed,
		Birthday: multiplyReserve.Birthday,
		Status: multiplyReserve.Status,
		Comment: strings.ToUpper(multiplyReserve.Comment),
		Notes: strings.ToUpper(multiplyReserve.Notes),
		Phone: multiplyReserve.Phone,
		PaymentMethod: multiplyReserve.PaymentMethod,
		PaymentDate: multiplyReserve.PaymentDate,
	}
}
func ValidateRouteChanges (route, ship, time string)error{
	if ship == "Gaviota" {

	}
	return nil
}