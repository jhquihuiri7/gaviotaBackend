package reports

import (
	"context"
	"fmt"
	"gaviotaBackend/variables"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
)

func GetDailyReportData(DailyRequest variables.DailyReportRequest) []variables.Reserve {
	var Reserves []variables.Reserve
	cursor, err := variables.ReservesGaviotaCollection.Find(context.TODO(), bson.D{{"time", DailyRequest.Time}, {"date", DailyRequest.Date}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var reserve variables.Reserve
		err = cursor.Decode(&reserve)
		if err != nil {
			log.Fatal(err)
		} else {
			Reserves = append(Reserves, reserve)
		}
	}

	return Reserves
}

func ApplyFontStyleTitle(cell, text string,f *excelize.File){
	f.SetCellRichText("Sheet1",cell,[]excelize.RichTextRun{
		{
			Text: text,
			Font: &excelize.Font{
				Size: 10,
				Bold: true,
			},
		}})
}
func ApplyFontStyleSubtitle(cell, text string,f *excelize.File){
	f.SetCellRichText("Sheet1",cell,[]excelize.RichTextRun{
		{
			Text: text,
			Font: &excelize.Font{
				Size: 7,
				Bold: true,
			},
		}})

}
func ApplyFontStyleText(cell, text string,f *excelize.File){
	f.SetCellRichText("Sheet1",cell,[]excelize.RichTextRun{
		{
			Text: text,
			Font: &excelize.Font{
				Size: 7,
			},
		}})

}
func ApplyFontStyleMini(cell, text1,text2,text3,text4 string,f *excelize.File){
	f.SetCellRichText("Sheet1",cell,[]excelize.RichTextRun{
		{
			Text: text1,
			Font: &excelize.Font{
				Size: 7,
				Bold: true,
			},
		},
		{
			Text: text2,
			Font: &excelize.Font{
				Size: 7,
			},
		},
		{
			Text: text3,
			Font: &excelize.Font{
				Size: 7,
				Bold: true,
			},
		},
		{
			Text: text4,
			Font: &excelize.Font{
				Size: 7,
			},
		},
	},
	)

}
func SetCellStyleMini(hcell, vcell string, f *excelize.File){
	f.MergeCell("Sheet1",hcell,vcell)
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "808080", Style: 2},
			{Type: "top", Color: "808080", Style: 2},
			{Type: "bottom", Color: "808080", Style: 2},
			{Type: "right", Color: "808080", Style: 2},
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "justify",
			Indent:          0,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     true,
			Vertical:        "center",
			WrapText:        true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle("Sheet1", hcell, vcell, style)
}
func SetCellStyleSubtitle(hcell, vcell string, f *excelize.File){
	f.MergeCell("Sheet1",hcell,vcell)
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "808080", Style: 2},
			{Type: "top", Color: "808080", Style: 2},
			{Type: "bottom", Color: "808080", Style: 2},
			{Type: "right", Color: "808080", Style: 2},
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          0,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     true,
			Vertical:        "center",
			WrapText:        true,
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#d5dce4"}, Pattern: 1},
	})
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle("Sheet1", hcell, vcell, style)
}
func SetCellStyleText(hcell, vcell string, f *excelize.File){
	f.MergeCell("Sheet1",hcell,vcell)
	style, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "808080", Style: 2},
			{Type: "top", Color: "808080", Style: 2},
			{Type: "bottom", Color: "808080", Style: 2},
			{Type: "right", Color: "808080", Style: 2},
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          0,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     true,
			Vertical:        "center",
			WrapText:        true,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetCellStyle("Sheet1", hcell, vcell, style)
}

func GenerateContent (reserves []variables.Reserve, f *excelize.File){
	var index int
	var indexString string
	for i, v := range reserves {
		if v.Age >= 2 {
			index++
			indexString = fmt.Sprintf("%d",index)
		}else {
			indexString = ""
		}
		if v.Age < 0 {
			v.Age = 0
		}
		res := ""
		tem := ""
		tur := ""
		switch v.Status {
		case "Residente":
			res ="x"
		case "Temporal":
			tem ="x"
		case "Turista":
			tur ="x"
		default:
			res ="x"
		}
		ApplyFontStyleText(fmt.Sprintf("C%d",14+i),indexString,f)
		ApplyFontStyleText(fmt.Sprintf("D%d",14+i),strings.ToUpper(v.Passenger),f)
		ApplyFontStyleText(fmt.Sprintf("H%d",14+i),strings.ToUpper(v.Passport),f)
		ApplyFontStyleText(fmt.Sprintf("K%d",14+i),strings.ToUpper(v.Country),f)
		ApplyFontStyleText(fmt.Sprintf("N%d",14+i),fmt.Sprintf("%d",v.Age),f)
		ApplyFontStyleText(fmt.Sprintf("P%d",14+i),res,f)
		ApplyFontStyleText(fmt.Sprintf("Q%d",14+i),tem,f)
		ApplyFontStyleText(fmt.Sprintf("R%d",14+i),tur,f)
		ApplyFontStyleText(fmt.Sprintf("S%d",14+i),v.Phone,f)
		ApplyFontStyleText(fmt.Sprintf("U%d",14+i),strings.ToUpper(v.Comment),f)

		SetCellStyleText(fmt.Sprintf("A%d",14+i),fmt.Sprintf("A%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("B%d",14+i),fmt.Sprintf("B%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("C%d",14+i),fmt.Sprintf("C%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("D%d",14+i),fmt.Sprintf("G%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("H%d",14+i),fmt.Sprintf("J%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("K%d",14+i),fmt.Sprintf("M%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("N%d",14+i),fmt.Sprintf("O%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("P%d",14+i),fmt.Sprintf("P%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("Q%d",14+i),fmt.Sprintf("Q%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("R%d",14+i),fmt.Sprintf("R%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("S%d",14+i),fmt.Sprintf("T%d",14+i),f)
		SetCellStyleText(fmt.Sprintf("U%d",14+i),fmt.Sprintf("W%d",14+i),f)
	}
}
