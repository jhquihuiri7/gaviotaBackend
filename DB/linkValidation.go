package DB

import (
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
	"time"
)

func GenerateLink(w http.ResponseWriter, r *http.Request){
	var response variables.RequestResponse
	newLink := variables.Link{
		LinkId:       strings.Split(uuid.NewV4().String(),".")[0],
		CreationDate: time.Now(),
	}
	response.Succes = newLink.LinkId
	variables.Links = append(variables.Links, newLink)
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
func ValidateLink(w http.ResponseWriter, r *http.Request){
	var response variables.RequestResponse
	params := mux.Vars(r)
	expire := params["linkToken"]

	for _, v := range variables.Links {
		if v.LinkId == expire {
			duration := time.Now().Sub(v.CreationDate)
			if duration < time.Hour*2 {
				response.Succes = "Link válido"
				response.Error = ""
			}else {
				response.Error = "Link no válido"
			}
			break
		}else {
			response.Error = "No se encontró link"
		}
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
