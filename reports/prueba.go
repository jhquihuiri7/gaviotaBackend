package reports

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"net/http"
	"os"
)

func PDF(w http.ResponseWriter, r *http.Request) {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(false)
	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(4, func() {
				m.Text("Jhonatan Q", props.Text{
					Top:         2,
					Size:        10,
					Extrapolate: true,
				})
			})
			m.ColSpace(4)
		})

	})

	m.Row(40, func() {
		m.Col(4, func() {
			m.Text("Gopher International Shipping, Inc.", props.Text{
				Top:         12,
				Size:        20,
				Extrapolate: true,
			})
			m.Text("1000 Shipping Gopher Golang TN 3691234 GopherLand (GL)", props.Text{
				Size: 12,
				Top:  22,
			})
		})
		m.ColSpace(4)
	})

	m.Line(10)

	m.SetBorder(true)
	left, _, rigth, _ := m.GetPageMargins()
	width, _ := m.GetPageSize()
	fmt.Println(width)
	fmt.Println(width - left - rigth)
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("Gaviota", props.Text{
				Align: consts.Center,
				Top:   10,
			})
		})
	})
	m.SetBorder(true)
	m.TableList([]string{"Hola", "Hola"}, [][]string{{"Perro", "GAtos"}, {"Perro", "GAtos"}}, props.TableList{
		HeaderProp: props.TableListContent{
			Size: 9,
		},
		ContentProp: props.TableListContent{
			Size: 8,
		},
		Align:                  consts.Center,
		VerticalContentPadding: 0,
		//AlternatedBackground: &grayColor,
		HeaderContentSpace: -0.25,
		Line:               false,
	})
	m.SetBorder(false)

	m.Row(40, func() {
		m.Col(4, func() {
			m.Text("João Sant'Ana 100 Main Street Stringfield TN 39021 United Stats (USA)", props.Text{
				Size: 15,
				Top:  12,
			})
		})
		m.ColSpace(4)
		m.Col(4, func() {
			m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
	})

	m.Line(10)

	m.Row(100, func() {
		m.Col(12, func() {
			_ = m.Barcode("https://github.com/johnfercher/maroto", props.Barcode{
				Center:  true,
				Percent: 70,
			})
			m.Text("https://github.com/johnfercher/maroto", props.Text{
				Size:  20,
				Align: consts.Center,
				Top:   65,
			})
		})
	})
	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(6, func() {
			m.Text("CODE: 123412351645231245564 DATE: 20-07-1994 20:20:33", props.Text{
				Size: 15,
				Top:  14,
			})
		})
		m.Col(6, func() {
			m.Text("CA", props.Text{
				Top:   1,
				Size:  85,
				Align: consts.Center,
			})
		})
	})

	dd, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	w.Header().Set("Content-Disposition", "attachment; filename=countries.pdf")
	w.Header().Set("Content-Type", r.Header.Get("application/pdf"))

	w.Write(dd.Bytes())
	//fmt.Println(dd.String())

}
