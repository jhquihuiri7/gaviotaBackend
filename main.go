package main

import (
	"fmt"
	"gaviotaBackend/DB"
	"gaviotaBackend/authentication"
	"gaviotaBackend/dev"
	"gaviotaBackend/middleware"
	"gaviotaBackend/payments"
	"gaviotaBackend/reports"
	"gaviotaBackend/reserves"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"gaviotaBackend/ws"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
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
	variables.BinReservesCollection = database.Collection("BinReserves")
	variables.PaymentsSystemHistory = database.Collection("PaymentsSystemHistory")
	authentication.InitKeys()
	ws.Connection = ws.WsConnection()
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
	router.HandleFunc("/api/getReservesDaily", reserves.GetReservesDaily).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReserveBase", reserves.EditReserveBase).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteReserve", reserves.DeleteReserve).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/randomReserves", utils.GetRandomReserves).Methods("GET", "OPTIONS")

	//reports
	router.HandleFunc("/api/dailyReport", reports.PdfDaily).Methods("GET", "OPTIONS")

	//payments
	router.HandleFunc("/response", payments.PaymentResponse).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getPaymentHistory", payments.GetPaymentHistory).Methods("GET", "OPTIONS")

	//dev
	router.HandleFunc("/admin/addAllUsers", dev.AddAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/addAllReferences", dev.AddAllReferences).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/getCountries", utils.CuntriesDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/downloadCountries", utils.DownloadCuntriesDB).Methods("GET", "OPTIONS")
	//tickets
	router.HandleFunc("/api/getPrinters", utils.GetPrinters).Methods("GET", "OPTIONS")

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
