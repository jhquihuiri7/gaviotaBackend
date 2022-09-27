package main

import (
	"fmt"
	"gaviotaBackend/DB"
	"gaviotaBackend/authentication"
	"gaviotaBackend/dev"
	"gaviotaBackend/middleware"
	"gaviotaBackend/reports"
	"gaviotaBackend/reserves"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"gaviotaBackend/ws"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
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
	authentication.InitKeys()

	tmp =  template.Must(template.ParseGlob("templates/*gohtml"))

}
var hub *ws.Hub
func main() {
	router = mux.NewRouter()


	router.Use(middleware.CORS)
	router.Use(middleware.JWT)
	router.HandleFunc("/", Index).Methods("GET", "OPTIONS")
	hub = ws.NewHub()
	go hub.Run()

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
	router.HandleFunc("/api/editReference", utils.EditReference).Methods("POST", "OPTIONS")

	//login
	router.HandleFunc("/api/login", authentication.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/validateToken", authentication.ValidateToken).Methods("GET", "OPTIONS")

	//reserves
	router.HandleFunc("/api/addReserves", reserves.AddReserves).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getReservesDaily", reserves.GetReservesDaily).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/editReserve",reserves.EditReserve).Methods("POST","OPTIONS")
	router.HandleFunc("/api/deleteReserve",reserves.DeleteReserve).Methods("POST","OPTIONS")
	router.HandleFunc("/api/randomReserves", utils.GetRandomReserves).Methods("GET", "OPTIONS")

	//dev
	router.HandleFunc("/admin/addAllUsers", dev.AddAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/addAllReferences", dev.AddAllReferences).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/getCountries", utils.CuntriesDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/downloadCountries", utils.DownloadCuntriesDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/admin/pdf", reports.PDF).Methods("GET", "OPTIONS")

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	router.HandleFunc("/ws",func (w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(conn.RemoteAddr())
	}).Methods("GET","OPTIONS")
	router.HandleFunc("/socket", Socket).Methods("GET","OPTIONS")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)

}
func Socket (w http.ResponseWriter, r *http.Request){
	tmp.Execute(w,nil)
}
func Index(w http.ResponseWriter, r *http.Request) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		fmt.Fprintln(w, tpl, met)
		return nil
	})
}

func wsEndpoint (w http.ResponseWriter, r *http.Request){
	fmt.Println("PASE A EJECUTAR EL ENDPOINT")
	r.Header.Add("Sec-Websocket-Version","13")
	ws.ServeWs(hub, w, r)
}
