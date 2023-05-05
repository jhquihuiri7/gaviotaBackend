package utils

import (
	"context"
	"fmt"
	"gaviotaBackend/overide/pdf"
	"gaviotaBackend/reports"
	"gaviotaBackend/tickets"
	"gaviotaBackend/variables"
	"github.com/sanketbajoria/maroto/pkg/consts"
	"github.com/sanketbajoria/maroto/pkg/props"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"
	"time"
)

type TicketLabel struct {
	Title, RouteLabel, Reference, Date, Departure, Route, Checkin, Total, Seller, SellDate, Conditions string
	ConditionLines                                                                                     []string
}

var SpanishLabel = TicketLabel{
	Title:      "Tickets de ferry, información & servicios turísticos",
	RouteLabel: "RUTA",
	Reference:  "Referencia",
	Date:       "Fecha",
	Departure:  "Hora",
	Route:      "Ruta",
	Checkin:    "Chequeo",
	Total:      "VALOR TOTAL",
	Seller:     "VENDEDOR",
	SellDate:   "FECHA DE VENTA",
	Conditions: "Condiciones del Viaje",
	ConditionLines: []string{
		"Cada  pasajero  puede  llevar 1 mochila de mano y 1 maleta de  20kg, si sobrepasa  el peso  tendrá  que  pagar  el  valor adicional ($10-$15).",
		"Si el pasajero no viaja debe informar a la agencia operadora donde compró su ticket con 24 horas de anticipación al viaje o perderá  su valor, con una penalidad  de  $5  por  cargos administrativos.",
		"Se recomienda  estar 50 minutos antes de la hora de salida.",
		"En caso de  retraso  del  pasajero  a  la  hora  de  salida,  la embarcación  no  se  responsabiliza por la pérdida del viaje.",
		"El servicio  portuario de taxis acuáticos ($1,00)  en cada isla y los impuestos  municipales en cada isla no están incluidos en el valor del ferry.",
	},
}
var EnglishLabel = TicketLabel{
	Title:      "Ferry tickets, information & tourist services",
	RouteLabel: "ROUTE",
	Reference:  "Reference",
	Date:       "Date",
	Departure:  "Time",
	Route:      "Route",
	Checkin:    "Check-in",
	Total:      "TOTAL",
	Seller:     "SELLER",
	SellDate:   "SALE DATE",
	Conditions: "Travel Conditions",
	ConditionLines: []string{
		"Each  passenger  can  carry 1 backpack  and  1 suitcase  of 20kg, if  it  exceeds  the  weight  you  will  have  to   pay  the  additional value ($10-$15).",
		"If   the   passenger  does   not  travel, they  must inform  the operating agency where they purchased their ticket 24 hours in advance of the trip or it will lose its value, with a penalty of $5 for administrative charges.",
		"It is recommended to be 50 minutes before the departure.",
		"In case of  delay  of the  passenger at the  time of departure, the boat is not responsible for the loss of the trip.",
		"The  port  service  of  water taxis ($1.00) on each island and the municipal  taxes  on  each  island are not included in the value of the ferry.",
	},
}

var getLang = map[string]TicketLabel{
	"es": SpanishLabel,
	"en": EnglishLabel,
}

func PrintTicket(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	reserve := param["reserve"][0]
	user := param["user"][0]
	lang := param["lang"][0]
	label := getLang[lang]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user", strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}
	ticketsData, total := tickets.GetReservesTicket2(reserve)
	litleSize := 8.0

	m := pdf.NewMarotoCustomSize(consts.Portrait, "c6", "mm", 78.0, getTickeHeight(len(ticketsData), "clientTicket"))

	m.SetBorder(false)
	m.SetPageMargins(0.5, 0, 1)

	m.SetBorder(false)

	m.Row(13, func() {
		m.FileImage(
			"./files/assets/logogaviotagrey.png",
			props.Rect{
				Percent: 100,
				Center:  true,
			})
	})
	m.Row(3, func() {
		m.Text(label.Title, props.Text{
			Align: consts.Center,
			Size:  litleSize,
			Style: consts.Bold,
		})
	})
	m.Row(3, func() {
		m.Text("Calle Española y Charles Darwin", props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(3, func() {
		m.Text("+593 99 892 7034 / +593 99 373 1079", props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(4, func() {
		m.Text("gaviota.ferry@gmail.com", props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	Routes(m, litleSize, ticketsData, lang)
	Total(m, litleSize, user, total, lang)
	m.Row(3, func() {})
	m.Row(3, func() {
		m.Text(label.Conditions, props.Text{Size: litleSize, Align: consts.Left, Style: consts.Bold})
	})
	m.Row(9, func() {
		m.Text(label.ConditionLines[0], props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(12, func() {
		m.Text(label.ConditionLines[1], props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(3, func() {
		m.Text(label.ConditionLines[2], props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(7, func() {
		m.Text(label.ConditionLines[3], props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(10, func() {
		m.Text(label.ConditionLines[4], props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Line(2)
	m.Row(3, func() {
		m.Text("(COPIA PARA OFICINA)", props.Text{Size: litleSize, Align: consts.Center})
	})
	Routes(m, litleSize, ticketsData, lang)
	Total(m, litleSize, user, total, lang)
	dd, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf")
	w.Write(dd.Bytes())
}
func Routes(m pdf.Maroto, litleSize float64, ticketsData []variables.Ticket, lang string) {
	label := getLang[lang]
	for i, v := range ticketsData {
		m.Row(2, func() {})
		m.Row(3, func() {
			m.Text(fmt.Sprintf("%s %d", label.RouteLabel, i+1), props.Text{Size: litleSize, Style: consts.BoldItalic, Align: consts.Center})
		})
		m.Row(2, func() {})
		m.TableList(
			[]string{"", ""},
			[][]string{{label.Reference, fmt.Sprintf("%s x%d", v.Passengers[0][0], len(v.Passengers))}},
			props.TableList{
				ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{8, 16}},
				HeaderContentSpace: -5,
				Line:               true,
				Align:              consts.Middle,
			})
		m.TableList(
			[]string{"", "", "", ""},
			[][]string{
				{label.Date, fmt.Sprintf("%s", v.Routes[0][0]), label.Departure, fmt.Sprintf("%s", v.Routes[0][3])},
				{label.Route, fmt.Sprintf("%s-%s", tickets.TranslateRoute(v.RouteId[0])[0], tickets.TranslateRoute(v.RouteId[0])[1]), label.Checkin, fmt.Sprintf("%s", v.Routes[0][2])},
				{"Ferry", fmt.Sprintf("%s", v.Routes[0][1]), "Paxs", fmt.Sprintf("%d", len(v.Passengers))}},
			props.TableList{
				ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{2, 14, 4, 4}},
				Line:               true,
				Align:              consts.Middle,
				HeaderContentSpace: -5,
			})
	}

}
func Total(m pdf.Maroto, litleSize float64, user string, total int, lang string) {
	label := getLang[lang]
	m.Row(3, func() {
		m.Text(fmt.Sprintf("%s: $%d", label.Total, total), props.Text{Size: litleSize})
	})
	m.Row(3, func() {
		m.Text(fmt.Sprintf("%s: %s", label.Seller, user), props.Text{Size: litleSize})
	})
	m.Row(3, func() {
		m.Text(fmt.Sprintf("%s: %s", label.SellDate, reports.FormatDate(primitive.NewDateTimeFromTime(time.Now()))[6:]), props.Text{Size: litleSize})
	})
}
func getTickeHeight(nRoutes int, ticketType string) float64 {
	switch ticketType {
	case "clientTicket":
		if nRoutes == 0 {
			return 105.0
		} else {
			return 93.0 + 24 + (51 * float64(nRoutes))
		}
	case "copyTicket":
		if nRoutes == 1 {
			return 80
		} else {
			return 25.0 + 12 + (21 * float64(nRoutes))
		}
	default:
		return 80.0
	}
}
