package utils

import (
	"encoding/json"
	"fmt"
	"gaviotaBackend/tickets"
	"gaviotaBackend/variables"
	"net/http"
	"strings"
)

func ShareOtherReserves (w http.ResponseWriter, r *http.Request){
	param := r.URL.Query()
	reserve := param["reserve"][0]
	ticketsData, _ := tickets.GetReservesTicketMini(reserve)
	message := "https://wa.me/?text="
	for _, v := range ticketsData {
		if v.Routes[0][1] != "GAVIOTA" {
			message += fmt.Sprintf("%s%s%s%s %s%d PAX%s%s %s%s %s-%s%s %s%s%s","%2A",v.Routes[0][0],"%2A","%0A","%2A",len(v.Passengers),"%2A","%0A",v.Routes[0][1],"%0A",tickets.TranslateRoute(v.RouteId[0])[0],tickets.TranslateRoute(v.RouteId[0])[1],"%0A",v.Routes[0][3],"%0A","%0A")
			for _, val := range v.Passengers {
				message += fmt.Sprintf("%s%s %s%s %s%s %s años%s %s%s%s", val[0],"%0A", val[1],"%0A",val[2],"%0A",val[3],"%0A",strings.ToUpper(val[4]),"%0A","%0A")
			}
		}
		message += "%0A"
	}
	message = strings.ReplaceAll(message," ","%20")
	response := variables.RequestResponse{Succes: message}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w,string(JSONresponse))
}
