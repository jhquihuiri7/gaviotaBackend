package main

import (
	"context"
	"fmt"
	"gaviotaBackend/DB"
	"gaviotaBackend/authentication"
	"gaviotaBackend/dev"
	"gaviotaBackend/middleware"
	"gaviotaBackend/more"
	"gaviotaBackend/payments"
	"gaviotaBackend/reports"
	"gaviotaBackend/reserves"
	"gaviotaBackend/sales"
	"gaviotaBackend/storage"
	"gaviotaBackend/tickets"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"gaviotaBackend/ws"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

var router *mux.Router
var tmp *template.Template

func init() {
	database := DB.ConnectDB()
	variables.UsersCollection = database.Collection("Users")
	variables.ReferencesCollection = database.Collection("References")
	variables.FrequentsCollection = database.Collection("Usuarios")
	variables.ReservesGaviotaCollection = database.Collection("ReservesGaviota")
	variables.ReservesOtherCollection = database.Collection("ReservesOther")
	variables.ReservesExternalCollection = database.Collection("ReservesExternal")
	variables.BinReservesCollection = database.Collection("BinReserves")
	variables.PaymentsSystemHistory = database.Collection("PaymentsSystemHistory")
	variables.PaymentsReferenceHistory = database.Collection("PaymentsReferenceHistory")
	variables.FrequentsNewCollection = database.Collection("Frequents")
	authentication.InitKeys()
	ws.Connection = ws.WsConnection()
	variables.CloudinaryStorage = storage.InitStorage()
	tmp = template.Must(template.ParseGlob("templates/*gohtml"))

}
func main() {
	router = mux.NewRouter()

	router.Use(middleware.CORS)
	router.Use(middleware.JWT)
	router.HandleFunc("/", Index).Methods("GET", "OPTIONS")

	//user
	router.HandleFunc("/api/addUser", utils.AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getUser", utils.GetUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getAllUsers", utils.GetUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/editUser", utils.EditUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteUser", utils.DeleteUser).Methods("POST", "OPTIONS")

	//reference
	router.HandleFunc("/api/addReference", utils.AddReference).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getAllReferences", utils.GetReference).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteReference", utils.DeleteReference).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/restoreReserve", reserves.RestoreBin).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReference", utils.EditReference).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getBin", reserves.GetBin).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cleanBin", reserves.CleanBin).Methods("GET", "OPTIONS")
	//login
	router.HandleFunc("/api/login", authentication.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/validateToken", authentication.ValidateToken).Methods("GET", "OPTIONS")

	//reserves
	router.HandleFunc("/api/addReserves", reserves.AddReserves).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/addReservesExternal", reserves.AddReservesExternal).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getReservesDaily", reserves.GetReservesDaily).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getReservesExternal", reserves.GetReservesExternal).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/editReserveBase", reserves.EditReserveBase).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReserveSingle", reserves.EditReserveSingle).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReserveExternal", reserves.EditReserveExternal).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReserveExternalBase", reserves.EditReserveExternalBase).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteReserve", reserves.DeleteReserve).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/randomReserves", utils.GetRandomReserves).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/generateLink",DB.GenerateLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/validateLink/expireAt/{linkToken}",DB.ValidateLink).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getRecentAddedReserves",reserves.GetRecentAddedReserves).Methods("GET", "OPTIONS")

	//utils
	router.HandleFunc("/api/ShareOtherReserves",utils.ShareOtherReserves).Methods("GET", "OPTIONS")
	//frequents
	router.HandleFunc("/api/getFrequents",utils.GetFrequents).Methods("GET","OPTIONS")

	//reports
	router.HandleFunc("/api/dailyReport", reports.DailyExcelReport).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getSalesReport",sales.SaleReport).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getSalesReportOther",sales.SaleReportOther).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getBalanceReport",sales.BalanceReport).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/registerAdvance",sales.RegisterAdvance).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getHistoryReport",sales.PaymentHistoryReport).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getDailyUserSales",more.GetDailyUserSales).Methods("POST", "OPTIONS")

	//payments
	router.HandleFunc("/response", payments.PaymentResponse).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getPaymentHistory", payments.GetPaymentHistory).Methods("GET", "OPTIONS")

	//dev
	router.HandleFunc("/admin/addAllUsers", dev.AddAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/addAllReferences", dev.AddAllReferences).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/getCountries", utils.CuntriesDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/downloadCountries", utils.DownloadCuntriesDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/getFrequents", dev.GetFrequents).Methods("GET", "OPTIONS")

	//tickets
	router.HandleFunc("/api/generateTicket",tickets.GenerateTicket1).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/generateTicket2",tickets.GenerateTicket2).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getPrinters", utils.GetPrinters).Methods("GET", "OPTIONS")


	go func(){
		for {
			result, err := variables.CloudinaryStorage.Admin.Assets(context.TODO(),admin.AssetsParams{AssetType: "raw"})
			if err != nil {
				fmt.Println(err)

			}else {
				var publicIds []string
				for _, v := range result.Assets {
					if strings.Contains(v.SecureURL,"GaviotaFerry/Reports"){
						fmt.Println(v.PublicID)
						publicIds = append(publicIds, v.PublicID)
					}

				}
				deleted, err :=variables.CloudinaryStorage.Admin.DeleteAssets(context.TODO(),admin.DeleteAssetsParams{PublicIDs: publicIds, AssetType: "raw"})
				if err != nil {
					fmt.Println(err)
				}else{
					fmt.Println(deleted.Deleted)
				}
			}
			//Dutation == 1 week
			time.Sleep(time.Hour *168)
		}
	}()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		data := variables.Reserve{Id: "fsdfsd", User: "Fsdfsd", Ship: "Gaviota"}
		var datas []interface{}
		datas = append(datas, data)
		ws.SendSocketMessage(datas)
	}).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)

}

func Index(w http.ResponseWriter, r *http.Request) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		fmt.Fprintln(w, tpl, met)
		return nil
	})
}
