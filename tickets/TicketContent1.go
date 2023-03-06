package tickets

import (
	"fmt"
	"gaviotaBackend/variables"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateTicketContent1(m pdf.Maroto, ticketData variables.Ticket, litleSize float64, showTotal bool, lang string) {
	var blueColor = color.Color{Red: 105, Green: 197, Blue: 197}
	var text []string

	if lang == "es" {
		text = append(text, "Fecha de viaje")
		text = append(text, "Lancha")
		text = append(text, "Hora de chequeo")
		text = append(text, "Salida")
		text = append(text, "Nombre del pasajero")
		text = append(text, "Nacionalidad")
		text = append(text, "Cédula")
		text = append(text, "Vendedor:")
		text = append(text, "TOTAL:")
	} else {
		text = append(text, "Travel date")
		text = append(text, "Ferry")
		text = append(text, "Check-in time")
		text = append(text, "Departure")
		text = append(text, "Name of the passenger")
		text = append(text, "Nationality")
		text = append(text, "Passport")
		text = append(text, "Seller:")
		text = append(text, "TOTAL:")
	}

	m.Row(2, func() {})
	m.TableList(
		[]string{text[0], text[1], text[2], text[3]},
		ticketData.Routes,
		props.TableList{
			HeaderProp: props.TableListContent{
				Color: blueColor,
			},
			Align: consts.Center,
			ContentProp: props.TableListContent{
				Size: litleSize,
			},
			HeaderContentSpace:     0,
			VerticalContentPadding: 0,
		},
	)
	m.Row(5, func() {})
	m.TableList(
		[]string{text[4], text[5]},
		ticketData.Passengers,
		props.TableList{
			Align: consts.Center,
			HeaderProp: props.TableListContent{
				Color: blueColor,
			},
			ContentProp: props.TableListContent{
				Size: litleSize,
			},
			HeaderContentSpace:     0,
			VerticalContentPadding: 0,
		},
	)

	if showTotal {
		m.Row(12, func() {
			m.Col(2, func() {
				m.Text(text[7])
			})
			m.Col(2, func() {
				m.Text("Moraima Freire")
			})
			m.Col(4, func() {
				m.Text("")
			})
			m.Col(2, func() {
				m.Text(text[8])
			})
			m.Col(2, func() {
				m.Text(fmt.Sprintf("$%d.00", ticketData.Total))
			})
		})
	}
	m.Row(5, func() {})
}
