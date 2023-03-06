package tickets

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateHeaderTicket(m pdf.Maroto, litleSize float64, lang string) {
	//m.Row(20, func() {
	//	m.Col(12, func() {
	//		m.Text("PASAJES/INFORMACIÓN Y SERVICIOS TURÍSTICOS", props.Text{
	//			Align: consts.Center,
	//			Size:  titleSize,
	//			Style: consts.Bold,
	//		})
	//	})
	//})
	m.Row(30, func() {
		m.Col(6, func() {
			m.FileImage(
				"./files/assets/lOGO-01.png",
				props.Rect{
					Percent: 100,
					Center:  true,
				},
			)
		})
		var text []string
		if lang == "es" {
			text = append(text, "TICKET DE LANCHA")
			text = append(text, "Av. Española y Charles Darwin")
			text = append(text, "Celular: 0993731079")
			text = append(text, "Correo: ventasdarwinscubadive@gmail.com")
		} else {
			text = append(text, "FERRY TICKET")
			text = append(text, "Av. Española y Charles Darwin")
			text = append(text, "Phone: 0993731079")
			text = append(text, "Email: ventasdarwinscubadive@gmail.com")
		}

		m.Col(6, func() {
			m.Text(text[0], props.Text{
				Align:       consts.Right,
				Size:        12,
				Style:       consts.Bold,
				Extrapolate: true,
			})
			m.Text(text[1], props.Text{
				Align: consts.Right,
				Size:  litleSize,
				Top:   7,
			})
			m.Text(text[2], props.Text{
				Align: consts.Right,
				Size:  litleSize,
				Top:   11,
			})
			m.Text(text[3], props.Text{
				Align: consts.Right,
				Size:  litleSize,
				Top:   15,
			})
		})
	})
}
