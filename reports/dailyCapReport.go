package reports

import (
	"fmt"
	"gaviotaBackend/variables"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func PdfDaily(w http.ResponseWriter, r *http.Request) {
	var DailyRequest variables.DailyReportRequest
	param := r.URL.Query()
	t := param["time"][0]
	date := strings.ReplaceAll(param["date"][0], " ", "+")
	nd, err := time.Parse("2006-01-02", date)
	dateTime := primitive.NewDateTimeFromTime(nd)
	DailyRequest.Time = t
	DailyRequest.Date = dateTime
	//err := json.NewDecoder(r.Body).Decode(&DailyRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DailyRequest.Date)

	//newPdf
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(25, 2.5, 25)

	m.RegisterHeader(func() {
		m.SetBorder(false)
		m.Row(20, func() {
			m.Col(4, func() {
				m.FileImage(
					"./files/assets/header-icon.png",
					props.Rect{
						Center:  true,
						Percent: 80,
					},
				)
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {

		})
	})
	m.SetBorder(true)
	m.Row(6, func() {
		m.Col(12, func() {
			m.Text("COE PROVINCIAL DE GALAPAGOS", props.Text{
				Align: consts.Center,
				Size:  12,
				Style: consts.Bold,
			})
		})
	})
	m.Row(5, func() {
		m.Col(12, func() {
			m.Text("FORMATO DE CONTROL DE PASAJEROS AEREO Y MARITIMO", props.Text{
				Align: consts.Center,
				Size:  10,
				Style: consts.Bold,
			})
		})
	})
	m.Row(4, func() {
		m.Col(12, func() {
			m.Text("NOMBRE DEL TRANSPORTE: LANCHA GAVIOTA", props.Text{
				Align: consts.Center,
				Size:  8,
				Style: consts.Bold,
			})
		})
	})
	m.Row(3, func() {
		m.Col(4, func() {
			m.Text(FormatRoute(DailyRequest.Time), props.Text{
				Align: consts.Center,
				Size:  6,
			})
		})
		m.Col(4, func() {
			m.Text(FormatDate(DailyRequest.Date), props.Text{
				Align: consts.Center,
				Size:  6,
			})
		})
		m.Col(4, func() {
			m.Text(FormatTime(DailyRequest.Time), props.Text{
				Align: consts.Center,
				Size:  6,
			})
		})
	})
	m.TableList([]string{"N°", "NOMBRES Y APELLIDOS", "PAIS", "CEDULA", "EDAD"}, GetDailyReportData(DailyRequest), props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 4, 3, 3, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 4, 3, 3, 1},
		},
		Align:                  consts.Center,
		VerticalContentPadding: 0,
		HeaderContentSpace:     -0.25,
		Line:                   false,
	})
	m.Row(15, func() {
		m.Col(12, func() {
			m.Text("CAP. PABLO VERA", props.Text{
				Align: consts.Center,
				Size:  12,
				Style: consts.Bold,
				Top:   10,
			})
		})
	})
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("LANCHA GAVIOTA", props.Text{
				Align: consts.Center,
				Size:  12,
				Style: consts.Bold,
			})
		})
	})

	dd, err := m.Output()

	if err != nil {
		fmt.Println("Could not save PdfDaily:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", FormatDailyName(DailyRequest.Time, DailyRequest.Date)))
	//w.Header().Set("Content-Type", r.Header.Get("application/pdf"))

	w.Write(dd.Bytes())
	//fmt.Fprintln(w,string(dd.Bytes()))

}
