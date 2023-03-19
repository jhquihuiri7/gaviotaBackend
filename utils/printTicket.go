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

func PrintTicket(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	reserve := param["reserve"][0]
	user := param["user"][0]
	lang := param["lang"][0]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user", strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}
	ticketsData, total := tickets.GetReservesTicket2(reserve)
	litleSize := 8.0
	var text []string
	if lang == "es" {
		text = append(text, "Tickets de ferry, información & servicios turísticos")
		text = append(text, "Calle Española y Charles Darwin")
		text = append(text, "+593 99 892 7034 / +593 99 373 1079")
		text = append(text, "gaviota.ferry@gmail.com")
	} else {
		text = append(text, "FERRY TICKET")
		text = append(text, "Av. Española y Charles Darwin")
		text = append(text, "Phone: 0993731079")
		text = append(text, "Email: gaviota.ferry@gmail.com")
	}
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
		m.Text(text[0], props.Text{
			Align: consts.Center,
			Size:  litleSize,
			Style: consts.Bold,
		})
	})
	m.Row(3, func() {
		m.Text(text[1], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(3, func() {
		m.Text(text[2], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(4, func() {
		m.Text(text[3], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	Routes(m, litleSize, ticketsData)
	Total(m, litleSize, user, total)
	m.Row(3, func() {})
	m.Row(3, func() {
		m.Text("Condiciones de viaje:", props.Text{Size: litleSize, Align: consts.Left, Style: consts.Bold})
	})
	m.Row(9, func() {
		m.Text("Cada  pasajero  puede  llevar 1 mochila de mano y 1 maleta de  20kg, si sobrepasa  el peso  tendrá  que  pagar  el  valor adicional ($10-$15).", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(12, func() {
		m.Text("Si el pasajero no viaja debe informar a la agencia operadora donde compró su ticket con 24 horas de anticipación al viaje o perderá  su valor, con una penalidad  de  $5  por  cargos administrativos.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(3, func() {
		m.Text("Se recomienda  estar 50 minutos antes de la hora de salida.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(7, func() {
		m.Text("En caso de  retraso  del  pasajero  a  la  hora  de  salida,  la embarcación  no  se  responsabiliza por la pérdida del viaje.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(10, func() {
		m.Text("El servicio  portuario de taxis acuáticos ($1,00)  en cada isla y los impuestos  municipales en cada isla no están incluidos en el valor del ferry.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Line(2)
	m.Row(3, func() {
		m.Text("(COPIA PARA OFICINA)", props.Text{Size: litleSize, Align: consts.Center})
	})
	Routes(m, litleSize, ticketsData)
	Total(m, litleSize, user, total)
	dd, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf")
	w.Write(dd.Bytes())
}
func PrintTicketCopy(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	reserve := param["reserve"][0]
	user := param["user"][0]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user", strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}
	ticketsData, total := tickets.GetReservesTicket2(reserve)
	litleSize := 8.0

	m := pdf.NewMarotoCustomSize(consts.Portrait, "A7", "mm", 80.0, getTickeHeight(len(ticketsData), "copyTicket"))
	m.SetPageMargins(1, 5, 1)

	m.Row(3, func() {
		m.Text("(COPIA PARA OFICINA)", props.Text{Size: litleSize, Align: consts.Center})
	})
	Routes(m, litleSize, ticketsData)
	Total(m, litleSize, user, total)
	dd, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf")
	w.Write(dd.Bytes())
}
func Routes(m pdf.Maroto, litleSize float64, ticketsData []variables.Ticket) {
	for i, v := range ticketsData {
		m.Row(2, func() {})
		m.Row(3, func() {
			m.Text(fmt.Sprintf("RUTA %d", i+1), props.Text{Size: litleSize, Style: consts.BoldItalic, Align: consts.Center})
		})
		m.Row(2, func() {})
		m.TableList(
			[]string{"", ""},
			[][]string{{"Referencia", fmt.Sprintf("%s x%d", v.Passengers[0][0], len(v.Passengers))}},
			props.TableList{
				ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{8, 16}},
				HeaderContentSpace: -5,
				Line:               true,
				Align:              consts.Middle,
			})
		m.TableList(
			[]string{"", "", "", ""},
			[][]string{
				{"Fecha", fmt.Sprintf("%s", v.Routes[0][0]), "Hora", fmt.Sprintf("%s", v.Routes[0][3])},
				{"Ruta", fmt.Sprintf("%s-%s", tickets.TranslateRoute(v.RouteId[0])[0], tickets.TranslateRoute(v.RouteId[0])[1]), "Chequeo", fmt.Sprintf("%s", v.Routes[0][2])},
				{"Ferry", fmt.Sprintf("%s", v.Routes[0][1]), "Paxs", fmt.Sprintf("%d", len(v.Passengers))}},
			props.TableList{
				ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{2, 14, 4, 4}},
				Line:               true,
				Align:              consts.Middle,
				HeaderContentSpace: -5,
			})
	}

}
func Total(m pdf.Maroto, litleSize float64, user string, total int) {
	m.Row(3, func() {
		m.Text(fmt.Sprintf("VALOR TOTAL: $%d", total), props.Text{Size: litleSize})
	})
	m.Row(3, func() {
		m.Text(fmt.Sprintf("VENDEDOR: %s", user), props.Text{Size: litleSize})
	})
	m.Row(3, func() {
		m.Text(fmt.Sprintf("FECHA DE VENTA: %v", reports.FormatDate(primitive.NewDateTimeFromTime(time.Now()))), props.Text{Size: litleSize})
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
