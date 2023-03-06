package utils

import (
	"context"
	"fmt"
	"gaviotaBackend/variables"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"strings"
)

func PrintTicket(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	//reserve := param["reserve"][0]
	user := param["user"][0]
	lang := param["lang"][0]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user", strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}
	//titleSize := 12.0
	litleSize := 5.0
	//conditionsSize := 7.0
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
	m := pdf.NewMarotoCustomSize(consts.Portrait, "A7", "mm", 74.0, 140.0)
	m.SetPageMargins(5, 5, 5)

	m.Row(10, func() {
		m.FileImage(
			"./files/assets/logogaviotagrey.png",
			props.Rect{
				Percent: 100,
				Center:  true,
			})
	})

	m.Row(2, func() {
		m.Text(text[0], props.Text{
			Align: consts.Center,
			Size:  6,
			Style: consts.Bold,
		})
	})
	m.Row(2, func() {
		m.Text(text[1], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(2, func() {
		m.Text(text[2], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	m.Row(5, func() {
		m.Text(text[3], props.Text{
			Align: consts.Center,
			Size:  litleSize,
		})
	})
	Routes(m, litleSize)
	m.Row(3, func() {})
	m.Row(3, func() {
		m.Text("Condiciones de viaje:", props.Text{Size: litleSize, Align: consts.Left, Style: consts.Bold})
	})
	m.Row(4, func() {
		m.Text("Cada pasajero puede llevar 1 mochila de mano y 1 maleta de 20kg, si excede el peso tendrá que pagar el valor adicional ($10-$15).", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(6, func() {
		m.Text("Si el pasajero no viaja deberá informar a la agencia operadora donde compró su ticket con 24 horas de anticipación o perderá su valor, con una penalidad  de $5 por cargos administrativos.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(2, func() {
		m.Text("Se recomienda estar 50 minutos antes de la hora de salida.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(4, func() {
		m.Text("En  caso  de  retraso  del   pasajero  a la hora de salida, la  embarcación  no  se  responsabiliza por la pérdida del viaje.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(4, func() {
		m.Text("El  servicio   portuario  de taxis  acuáticos ($1,00)  en cada  isla y los  impuestos municipales en cada isla no están incluidos en el valor del ferry.", props.Text{Size: litleSize, Align: consts.Left})
	})
	m.Row(3, func() {})
	m.Line(1)
	m.Row(2, func() {})
	m.Row(2, func() {
		m.Text("(COPIA PARA OFICINA)", props.Text{Size: litleSize, Align: consts.Center})
	})

	Routes(m, litleSize)
	dd, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf")
	w.Write(dd.Bytes())
}
func Routes(m pdf.Maroto, litleSize float64) {
	m.Row(3, func() {
		m.Text("RUTA 1", props.Text{Size: litleSize, Style: consts.Italic, Align: consts.Center})
	})
	m.TableList(
		[]string{"", ""},
		[][]string{{"Referencia", "Morayma Freire"}},
		props.TableList{
			ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{4, 8}},
			HeaderContentSpace: -5,
			Line:               true,
			Align:              consts.Middle,
		})
	m.TableList(
		[]string{"", "", "", ""},
		[][]string{
			{"Fecha", "12 de marzo de 2023", "Hora", "7:00 AM"},
			{"Ruta", "San Cristóbal - Santa Cruz", "Chequeo", "6:15 AM"},
			{"FERRY", "GAVIOTA", "Paxs", "4"}},
		props.TableList{
			ContentProp:        props.TableListContent{Size: litleSize, GridSizes: []uint{2, 5, 3, 2}},
			Line:               true,
			Align:              consts.Middle,
			HeaderContentSpace: -5,
		})
	m.Row(2, func() {
		m.Text("VALOR TOTAL: $240,00", props.Text{Size: litleSize})
	})
	m.Row(2, func() {
		m.Text("VENDEDOR: Mory Freire", props.Text{Size: litleSize})
	})
	m.Row(2, func() {
		m.Text("FECHA DE VENTA: 10-marzo", props.Text{Size: litleSize})
	})
}
