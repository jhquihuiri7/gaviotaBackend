package tickets

import (
	"fmt"
	"gaviotaBackend/variables"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateTicketContent2(m pdf.Maroto, ticketsData []variables.Ticket, litleSize float64, lang string) {
	var blueColor = color.Color{Red: 105, Green: 197, Blue: 197}
	var whiteColor = color.Color{Red: 255, Green: 255, Blue: 255}
	var text []string
	if lang == "es" {
		text = append(text, "RUTA")
		text = append(text, "Número de pasajeros:")
		text = append(text, "Isla de partida")
		text = append(text, "Isla de destino")
	} else {
		text = append(text, "ROUTE")
		text = append(text, "Number of passengers:")
		text = append(text, "Departure island")
		text = append(text, "Destination island")
	}
	for i, v := range ticketsData {

		m.Row(15, func() {
			//m.SetBackgroundColor(blueColor)
			m.Col(5, func() {
				m.Text(fmt.Sprintf("%s %d", text[0], i+1), props.Text{
					Align: consts.Center,
					Size:  15,
					Top:   4,
					Style: consts.Bold,
					Color: blueColor,
				})
			})
			m.Col(4, func() {
				m.Text(text[1], props.Text{
					Align: consts.Center,
					Size:  15,
					Top:   4,
					Style: consts.Bold,
					Color: blueColor,
				})
			})
			m.Col(3, func() {
				m.Text(fmt.Sprintf("%d", len(v.Passengers)), props.Text{
					Align: consts.Center,
					Size:  15,
					Top:   4,
					Style: consts.Bold,
				})
			})
		})

		m.Row(15, func() {
			m.SetBackgroundColor(blueColor)
			m.Col(5, func() {
				m.Text(text[2], props.Text{
					Size:  10,
					Color: whiteColor,
					Style: consts.Bold,
					Top:   2,
					Align: consts.Center,
				})
				m.Text(TranslateRoute(v.RouteId[0])[0], props.Text{
					Size:  15,
					Color: whiteColor,
					Style: consts.Bold,
					Top:   7,
					Align: consts.Center,
				})
			})
			m.Col(2, func() {
				m.FileImage("./files/assets/arrow.png", props.Rect{
					Percent: 60,
					Center:  true,
				})
			})
			m.Col(5, func() {
				m.Text(text[3], props.Text{
					Size:  10,
					Color: whiteColor,
					Style: consts.Bold,
					Top:   2,
					Align: consts.Center,
				})
				m.Text(TranslateRoute(v.RouteId[0])[1], props.Text{
					Size:  15,
					Color: whiteColor,
					Style: consts.Bold,
					Top:   7,
					Align: consts.Center,
				})
			})
		})
		m.SetBackgroundColor(whiteColor)
		GenerateTicketContent1(m, v, litleSize, false, lang)
	}
}
