package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"gaviotaBackend/DB"
	"gaviotaBackend/authentication"
	"gaviotaBackend/dev"
	"gaviotaBackend/middleware"
	"gaviotaBackend/reports"
	"gaviotaBackend/utils"
	"gaviotaBackend/variables"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
)

var router *mux.Router

func init() {
	database := DB.ConnectDB()
	variables.UsersCollection = database.Collection("Users")
	variables.ReferencesCollection = database.Collection("References")
	variables.FrequentsCollection = database.Collection("Usuarios")
	authentication.InitKeys()

}
func main() {
	router = mux.NewRouter()
	router.Use(middleware.JWT)
	router.Use(middleware.CORS)
	router.HandleFunc("/", Index).Methods("GET")
	//user
	router.HandleFunc("/api/addUser", utils.AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getAllUsers", utils.GetUser).Methods("GET")
	router.HandleFunc("/api/deleteUser", utils.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/updatePasswordUser", utils.ChangePasswordUser).Methods("POST", "OPTIONS")
	//reference
	router.HandleFunc("/api/addReference", utils.AddReference).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getAllReferences", utils.GetReference).Methods("GET")
	router.HandleFunc("/api/deleteReference", utils.DeleteReference).Methods("DELETE")
	router.HandleFunc("/api/editReference", utils.EditReference).Methods("POST", "OPTIONS")

	//login
	router.HandleFunc("/api/login", authentication.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/validateToken", authentication.ValidateToken).Methods("GET")
	//dev
	router.HandleFunc("/admin/addAllUsers", dev.AddAllUsers).Methods("GET")
	router.HandleFunc("/admin/addAllReferences", dev.AddAllReferences).Methods("GET")
	router.HandleFunc("/admin/getCountries", utils.CuntriesDB).Methods("GET")
	router.HandleFunc("/admin/downloadCountries", utils.DownloadCuntriesDB).Methods("GET")
	router.HandleFunc("/admin/pdf", reports.PDF).Methods("GET")
	router.HandleFunc("/admin/prueba", getFrequent).Methods("GET")

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

type Frequents struct {
	Id           string `bson:"_id"`
	Referencia   string `bson:"referencia"`
	Cedula       string `bson:"cedula"`
	Telefono     string `bson:"telefono"`
	Status       string `bson:"status"`
	Nacionalidad string `json:"nacionalidad"`
	Edad         int    `json:"edad"`
	DateRegister int64  `json:"dateRegister"`
}

func getFrequent(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create("files/Frequents.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)

	cursor, err := variables.FrequentsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalln(err)
	}
	err = writer.Write([]string{"name", "status", "country", "id", "phone", "age"})
	for cursor.Next(context.TODO()) {
		var freq Frequents
		err = cursor.Decode(&freq)
		if err != nil {
			log.Fatal(err)
		}
		err = writer.Write([]string{freq.Referencia, freq.Status, freq.Nacionalidad, freq.Cedula, freq.Telefono, fmt.Sprintf("%d", freq.Edad)})
		if err != nil {
			log.Fatal(err)
		}
	}
}
