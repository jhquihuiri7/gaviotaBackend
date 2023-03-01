package reports

import (
	"encoding/json"
	"fmt"
	"gaviotaBackend/storage"
	"gaviotaBackend/variables"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"time"
)

func DailyExcelReport(w http.ResponseWriter, r *http.Request) {
	var DailyRequest variables.DailyReportRequest
	var response variables.RequestResponse
	param := r.URL.Query()
	t := param["time"][0]

	date := strings.ReplaceAll(param["date"][0], " ", "+")
	date = strings.ReplaceAll(param["date"][0], "/", "-")
	if len(date) == 9 {
		date = "0" + date
	}
	if len(date) == 8 {
		date = "0" + date[:2] + "0" + date[2:]
	}

	nd, _ := time.Parse("01-02-2006", date)
	fmt.Println(date)
	fmt.Println(nd)
	dateTime := primitive.NewDateTimeFromTime(nd)
	DailyRequest.Time = t
	DailyRequest.Date = dateTime
	reserves := GetDailyReportData(DailyRequest)
	sheet := "Sheet1"
	f := excelize.NewFile()
	f.SetPageLayout(sheet, excelize.PageLayoutOrientation("portrait"))
	f.SetPageMargins(sheet, excelize.PageMarginBottom(0), excelize.PageMarginFooter(0), excelize.PageMarginHeader(0), excelize.PageMarginLeft(0.9), excelize.PageMarginRight(0.3), excelize.PageMarginTop(0.4))
	f.SetColWidth(sheet, "A", "Z", 4.5)
	f.SetColWidth(sheet, "A", "C", 3.5)
	f.SetColWidth(sheet, "D", "D", 7.5)
	f.SetColWidth(sheet, "H", "H", 2)
	f.SetColWidth(sheet, "M", "M", 2.5)
	f.SetColWidth(sheet, "N", "N", 4)
	f.SetColWidth(sheet, "T", "T", 3.5)
	f.SetColWidth(sheet, "U", "U", 2.5)
	f.SetColWidth(sheet, "W", "W", 3)

	for i := 1; i <= 65; i++ {
		f.SetRowHeight("Sheet1", i, 12)
		if i == 2 {
			f.SetRowHeight("Sheet1", i, 24)
		}
	}

	//f.SetRowHeight("Sheet1", 7, 22)
	HeaderSection(f)
	InformationSection(f, sheet, nd.String(), t)
	ContentSection(f, sheet, reserves)
	FooterSection(13+len(reserves), sheet, f)

	//file, err := ioutil.TempFile(os.TempDir(),fmt.Sprintf("ReporteDiario %s-*.xlsx",nd.String()[:10]))
	//defer os.Remove(file.Name())

	res := storage.UploadAsset(variables.CloudinaryStorage, f, nd)
	//m := gomail.NewMessage()
	//m.SetHeader("From", "jhonatan.quihuiri@gmail.com")
	////m.SetHeader("To", "marthacasti2103@gmail.com","renanerling@gmail.com","morycat1995@gmail.com","gaviota.ferry@gmail.com")
	//m.SetHeader("To", "jhonatan.quihuiri@gmail.com")
	////
	//m.SetHeader("Subject", fmt.Sprintf("REPORTE DIARIO %s",nd.String()[:10]))
	//m.SetBody("text/html", "")
	//m.Attach(file.Name())

	//d := gomail.NewPlainDialer("smtp.gmail.com", 587, "jhonatan.quihuiri@gmail.com", "dfouvrynuvwkydpi")
	//if err = d.DialAndSend(m); err != nil {
	//	response.Error = "Error al generar reporte"
	//}else {
	//	response.Succes = "Reporte enviado al correo electrónico"
	//}

	if strings.Contains(res, "clou") {
		response.Succes = fmt.Sprintf("https://view.officeapps.live.com/op/view.aspx?src=%s&wdOrigin=BROWSELINK", res)
	} else {
		response.Error = res
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
