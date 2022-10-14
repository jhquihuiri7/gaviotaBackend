package reserves

import "gaviotaBackend/variables"

func MultiplyReserve (multiplyReserve variables.MultiplyReserve)variables.Reserve{
	return variables.Reserve{
		Id:            multiplyReserve.Id,
		ReserveNumber: multiplyReserve.ReserveNumber,
		Passenger:     multiplyReserve.Passenger,
		Referece:      multiplyReserve.Referece,
		User:          multiplyReserve.User,
		Age:           multiplyReserve.Age,
		Date:          multiplyReserve.Date,
		Time:          multiplyReserve.Time,
		Passport:      multiplyReserve.Passport,
		Country:       multiplyReserve.Country,
		Price:         multiplyReserve.Price,
		Ship:          multiplyReserve.Ship,
		Route:multiplyReserve.Route,
		IsConfirmed:multiplyReserve.IsConfirmed,
		IsPayed:multiplyReserve.IsPayed,
		IsBlocked :multiplyReserve.IsBlocked,
	}
}
