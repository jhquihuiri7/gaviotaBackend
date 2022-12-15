package tickets

import (
	"context"
	"fmt"
	"gaviotaBackend/variables"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
	"strings"
)

func GenerateTicket1(w http.ResponseWriter, r *http.Request){
	var blueColor = color.Color{Red: 105, Green: 197,Blue: 197}
	var whiteColor = color.Color{Red: 255, Green: 255,Blue: 255}
	param := r.URL.Query()
	reserve := param["reserve"][0]
	user := param["user"][0]
	lang := param["lang"][0]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user",strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}

	titleSize := 12.0
	litleSize := 8.0
	//conditionsSize := 7.0
	var text []string
	if lang == "es" {
		text = append(text, "RUTA")
		text = append(text, "Número de pasajeros:")
		text = append(text, "Isla de partida")
		text = append(text, "Isla de destino")
	}else {
		text = append(text, "ROUTE")
		text = append(text, "Number of passengers:")
		text = append(text, "Departure island")
		text = append(text, "Destination island")
	}
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(15, 15, 15)

	GenerateHeaderTicket(m,litleSize,lang)
	m.Row(15, func(){
		//m.SetBackgroundColor(blueColor)
		m.Col(5,func(){
			m.Text(fmt.Sprintf("%s",text[0]), props.Text{
				Align: consts.Center,
				Size:  15,
				Top: 4,
				Style: consts.Bold,
				Color: blueColor,
			})
		})
		m.Col(4,func(){
			m.Text(text[1], props.Text{
				Align: consts.Center,
				Size:  15,
				Top: 4,
				Style: consts.Bold,
				Color: blueColor,
			})
		})
		m.Col(3,func(){
			m.Text(fmt.Sprintf("%d", 1), props.Text{
				Align: consts.Center,
				Size:  15,
				Top: 4,
				Style: consts.Bold,
			})
		})
	})
	m.Row(15,func(){
		m.SetBackgroundColor(blueColor)
		m.Col(5,func(){
			m.Text(text[2],props.Text{
				Size: 10,
				Color: whiteColor,
				Style: consts.Bold,
				Top: 2,
				Align: consts.Center,
			})
			m.Text("TranslateRoute(v.RouteId[0])[0]",props.Text{
				Size: 15,
				Color: whiteColor,
				Style: consts.Bold,
				Top: 7,
				Align: consts.Center,
			})
		})
		m.Col(2,func(){
			m.FileImage("./files/assets/arrow.png",props.Rect{
				Percent: 60,
				Center: true,
			})
		})
		m.Col(5,func(){
			m.Text(text[3],props.Text{
				Size: 10,
				Color: whiteColor,
				Style: consts.Bold,
				Top: 2,
				Align: consts.Center,
			})
			m.Text("TranslateRoute(v.RouteId[0])[1]",props.Text{
				Size: 15,
				Color: whiteColor,
				Style: consts.Bold,
				Top: 7,
				Align: consts.Center,
			})
		})
	})
	GenerateTicketContent1(m,GetReservesTicket(reserve), litleSize, true, lang)
	GenerateTicketCondition(m,titleSize,lang)
	dd, err := m.Output()

	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf" )
	w.Write(dd.Bytes())
}
func GenerateTicket2(w http.ResponseWriter, r *http.Request){
	param := r.URL.Query()
	reserve := param["reserve"][0]
	user := param["user"][0]
	lang := param["lang"][0]
	var userData variables.UsersData
	result := variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user",strings.ToUpper(user)}})
	err := result.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}
	ticketsData, total := GetReservesTicket2(reserve)

	titleSize := 12.0
	litleSize := 8.0
	//conditionsSize := 7.0
	var blueColor = color.Color{Red: 105, Green: 197,Blue: 197}

	var text []string

	if lang == "es" {
		text = append(text, "Vendedor:")
		text = append(text, "VALOR:")
	}else {
		text = append(text, "Seller:")
		text = append(text, "VALUE:")
	}
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(15, 15, 15)
	GenerateHeaderTicket(m,litleSize,lang)
	GenerateTicketContent2(m,ticketsData, litleSize,lang)
	m.Row(12,func(){
		m.Col(2,func(){
			m.Text(text[0],props.Text{Color: blueColor,Align: consts.Right})
		})
		m.Col(2,func(){
			m.Text(fmt.Sprintf("%s %s",userData.Name,userData.LastName),props.Text{Align: consts.Center})
		})
		m.Col(4,func(){
			m.Text("")
		})
		m.Col(2,func(){
			m.Text(text[1],props.Text{Color: blueColor,Align: consts.Right})
		})
		m.Col(2,func(){
			m.Text(fmt.Sprintf("$%d.00",total),props.Text{Align: consts.Center,Style: consts.Bold})
		})
	})
	GenerateTicketCondition(m,titleSize,lang)
	dd, err := m.Output()

	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=ticket.pdf" )
	w.Write(dd.Bytes())
}
