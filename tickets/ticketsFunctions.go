package tickets

import (
	"context"
	"fmt"
	"gaviotaBackend/reports"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sort"
	"strings"
)

func GetReservesTicket(reserveNumber string) variables.Ticket {
	var Reserves []variables.Reserve
	var total int
	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"reserve", reserveNumber}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			total += reserve.Price
			Reserves = append(Reserves, reserve)
		}
	}
	var passengers [][]string
	for _, v := range GetPassenger(Reserves) {
		passenger := []string{v.Passenger, v.Country}
		passengers = append(passengers, passenger)
	}
	var routes [][]string
	for _, v := range GetRoutes(Reserves) {
		route := []string{reports.FormatDate(v.Date)[7:], v.Ship, TranslateTimeCheckIn(v.Route, v.Time)[0], TranslateTimeCheckIn(v.Route, v.Time)[1]}
		routes = append(routes, route)
	}
	return variables.Ticket{Routes: routes, Passengers: passengers, Total: total}
}
func GetRoutes(reserves []variables.Reserve) []variables.TicketRoute {
	dates := make(map[primitive.DateTime]string)
	for _, v := range reserves {
		dates[v.Date] = v.Route
	}
	var ticketRoutes []variables.TicketRoute
	for i, _ := range dates {
		routes := make(map[string]string)
		for _, re := range reserves {
			if re.Date == i {
				routes[re.Route] = re.Time

			}
		}
		for index, ro := range routes {
			ticketRoute := variables.TicketRoute{Date: i, Route: index, Time: ro}
			ticketRoutes = append(ticketRoutes, ticketRoute)
		}
	}

	for i, v := range ticketRoutes {
		for _, re := range reserves {
			if v.Date == re.Date && v.Route == re.Route && v.Time == re.Time {
				ticketRoutes[i].Ship = re.Ship
			}
		}
	}
	sort.Slice(ticketRoutes, func(i, j int) bool {
		return ticketRoutes[i].Time < ticketRoutes[j].Time
	})
	sort.Slice(ticketRoutes, func(i, j int) bool {
		return ticketRoutes[i].Date < ticketRoutes[j].Date
	})
	return ticketRoutes
}
func GetPassenger(reserves []variables.Reserve) []variables.TicketPassenger {
	passengers := make(map[string]string)
	var ticketPassengers []variables.TicketPassenger
	for _, v := range reserves {
		passengers[v.Passenger] = v.Country
	}
	for i, v := range passengers {
		var ticketPassenger variables.TicketPassenger
		ticketPassenger.Passenger = i
		ticketPassenger.Country = v
		ticketPassengers = append(ticketPassengers, ticketPassenger)
	}
	return ticketPassengers
}

func GetReservesTicket2(reserveNumber string) ([]variables.Ticket, int) {
	var Reserves []variables.Reserve
	var total int
	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"reserve", reserveNumber}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			total += reserve.Price
			Reserves = append(Reserves, reserve)
		}
	}
	cursor, err = variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"reserve", reserveNumber}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			total += reserve.Price
			Reserves = append(Reserves, reserve)
		}
	}
	routes := GetRoutes(Reserves)
	var tickets []variables.Ticket
	for _, v := range routes {
		ticket := variables.Ticket{Routes: [][]string{{reports.FormatDate(v.Date)[7:], strings.ToUpper(v.Ship), TranslateTimeCheckIn(v.Route, v.Time)[0], TranslateTimeCheckIn(v.Route, v.Time)[1]}}, RouteId: []string{v.Route}}
		for _, re := range Reserves {
			if re.Date == v.Date && re.Route == v.Route && re.Time == v.Time && re.Ship == v.Ship {
				ticket.Passengers = append(ticket.Passengers, []string{re.Passenger, re.Country})
			}
		}
		tickets = append(tickets, ticket)
	}

	return tickets, total
}
func GetReservesTicketMini(reserveNumber string) ([]variables.Ticket, int) {
	var Reserves []variables.Reserve
	var total int
	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"reserve", reserveNumber}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			total += reserve.Price
			Reserves = append(Reserves, reserve)
		}
	}
	cursor, err = variables.ReservesOtherCollection.Find(context.TODO(), bson.D{{"reserve", reserveNumber}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			total += reserve.Price
			Reserves = append(Reserves, reserve)
		}
	}
	routes := GetRoutes(Reserves)
	var tickets []variables.Ticket
	for _, v := range routes {
		ticket := variables.Ticket{Routes: [][]string{{reports.FormatDate(v.Date)[7:], strings.ToUpper(v.Ship), TranslateTimeCheckIn(v.Route, v.Time)[0], TranslateTimeCheckIn(v.Route, v.Time)[1]}}, RouteId: []string{v.Route}}
		for _, re := range Reserves {
			if re.Date == v.Date && re.Route == v.Route && re.Time == v.Time && re.Ship == v.Ship {
				ticket.Passengers = append(ticket.Passengers, []string{re.Passenger, re.Country, re.Passport, fmt.Sprintf("%d", re.Age), re.Status})
			}
		}
		tickets = append(tickets, ticket)
	}

	return tickets, total
}
func TranslateRoute(route string) []string {
	translatedMap := make(map[string][]string)
	translatedMap["SX-SC"] = []string{"SANTA CRUZ", "SAN CRISTÓBAL"}
	translatedMap["SC-SX"] = []string{"SAN CRISTÓBAL", "SANTA CRUZ"}
	translatedMap["SX-IB"] = []string{"SANTA CRUZ", "ISABELA"}
	translatedMap["IB-SX"] = []string{"ISABELA", "SANTA CRUZ"}
	translatedMap["SX-FL"] = []string{"SANTA CRUZ", "FLOREANA"}
	translatedMap["FL-SX"] = []string{"FLOREANA", "SANTA CRUZ"}
	return translatedMap[route]
}
func TranslateTimeCheckIn(route, time string) []string {
	var translatedMap []string
	if route == "SX-SC" {
		if time == "Am" {
			translatedMap = append(translatedMap, "06:10")
			translatedMap = append(translatedMap, "07:00")
		} else {
			translatedMap = append(translatedMap, "14:10")
			translatedMap = append(translatedMap, "15:00")
		}
	} else if route == "SC-SX" {
		if time == "Am" {
			translatedMap = append(translatedMap, "06:10")
			translatedMap = append(translatedMap, "07:00")
		} else {
			translatedMap = append(translatedMap, "14:10")
			translatedMap = append(translatedMap, "15:00")
		}
	} else if route == "SX-IB" {
		if time == "Am" {
			translatedMap = append(translatedMap, "06:10")
			translatedMap = append(translatedMap, "07:00")
		} else {
			translatedMap = append(translatedMap, "14:10")
			translatedMap = append(translatedMap, "15:00")
		}
	} else if route == "IB-SX" {
		if time == "Am" {
			translatedMap = append(translatedMap, "05:20")
			translatedMap = append(translatedMap, "06:00")
		} else {
			translatedMap = append(translatedMap, "14:10")
			translatedMap = append(translatedMap, "15:00")
		}
	} else if route == "SX-FL" {
		translatedMap = append(translatedMap, "07:20")
		translatedMap = append(translatedMap, "08:00")
	} else if route == "FL-SX" {
		translatedMap = append(translatedMap, "14:10")
		translatedMap = append(translatedMap, "15:00")
	} else {
		if time == "Am" {
			translatedMap = append(translatedMap, "06:10")
			translatedMap = append(translatedMap, "07:00")
		} else {
			translatedMap = append(translatedMap, "14:10")
			translatedMap = append(translatedMap, "15:00")
		}
	}
	return translatedMap
}
