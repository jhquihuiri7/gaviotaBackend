package tickets

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateTicketCondition (m pdf.Maroto, titleSize float64, lang string){
	var orangeColor = color.Color{Red: 249, Green: 73,Blue: 3}
	var blueColor = color.Color{Red: 105, Green: 197,Blue: 197}
	var text []string
	if lang == "es" {
		text = append(text, "CONDICIONES DE VIAJE")
		text = append(text, "Un artículo personal")
		text = append(text, "de hasta 5 kg")
		text = append(text, "Una maleta en bodega")
		text = append(text, "de hasta 20 kg")
		text = append(text, "Servicio de taxis acuáticos ($1,00)")
		text = append(text, "no están incluidos en el valor del ferry")
		text = append(text, "Si el pasajero no viaja deberá informar a la agencia operadora donde compró su ticket con 48 horas de anticipación o perderá su valor. Con penalidad de $5 por cargos administrativos")
		text = append(text, "Se recomienda estar 40 minutos antes de la hora de salida.")
		text = append(text, "En caso de retraso del pasajero a la hora de salida, la embarcación no se responsabiliza por la pérdida del viaje.")
	}else {
		text = append(text, "TRAVEL CONDITIONS")
		text = append(text, "One personal item")
		text = append(text, "up to 5 kg")
		text = append(text, "A suitcase in hold")
		text = append(text, "up to 20kg")
		text = append(text, "Water taxi port service ($1.00)")
		text = append(text, "are not included in the value of the ferry")
		text = append(text, "If the passenger does not travel, they must inform to the travel agency where they bought their tickets 48 hours in advance or it will lose its value, with a penalty of $5 for administrative charges.")
		text = append(text, "It is recommended to be 45 minutes before the departure time of each ferry.")
		text = append(text, "In case of delay of the passenger at the time of departure, the ferry is not responsible for the loss of the trip.")
	}



	m.Row(10, func() {
		m.Col(12, func(){
			m.Text(text[0], props.Text{
				Align: consts.Left,
				Size:  titleSize,
				Style: consts.Bold,
				Color: orangeColor,
			})
		})
	})
	m.Row(50, func() {
		m.Col(2,func(){
			m.FileImage("./files/assets/flat.png",props.Rect{Percent: 70, Center: true})
		})
		m.Col(2,func(){
			m.Text(text[1], fancyConditionStyle(blueColor,10,15))
			m.Text(text[2], fancyConditionStyle(orangeColor,10,23))
		})
		m.Col(2,func(){
			m.FileImage("./files/assets/mochila.png", props.Rect{Percent: 100, Center: true})
		})
		m.Col(2,func(){
			m.Text(text[3], fancyConditionStyle(blueColor,10,15))
			m.Text(text[4], fancyConditionStyle(orangeColor,10,23))
		})
		m.Col(2,func(){
			m.FileImage("./files/assets/Dolar.png", props.Rect{Percent: 60, Center: true, Top: -50})
		})
		m.Col(2,func(){
			m.Text(text[5], fancyConditionStyle(blueColor,10,15))
			m.Text(text[6], fancyConditionStyle(orangeColor,10,22))
		})
	})

	m.Row(6, func() {
		m.Col(12, func(){
			m.Text(text[7], littleConditionStyle(orangeColor,8))
		})
	})
	m.Row(3, func() {
		m.Col(12, func(){
			m.Text(text[8], littleConditionStyle(orangeColor,8))
		})
	})
	m.Row(3, func() {
		m.Col(12, func(){
			m.Text(text[9], littleConditionStyle(orangeColor,8))
		})
	})
}
func littleConditionStyle(color color.Color, size float64) props.Text {
	return props.Text{
		Align: consts.Left,
		Size:  size,
		Style: consts.Normal,
		Color: color,
	}
}
func fancyConditionStyle(color color.Color, size float64, top float64) props.Text {
	return props.Text{
		Align: consts.Left,
		Top: top,
		Size:  size,
		Style: consts.Bold,
		Color: color,
	}
}