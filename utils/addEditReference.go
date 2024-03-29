package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

func AddReference(w http.ResponseWriter, r *http.Request) {
	references := variables.ReferencesData{}
	json.NewDecoder(r.Body).Decode(&references)
	references.Id = GenerateID()
	_, err := variables.ReferencesCollection.InsertOne(context.TODO(), references)
	response := variables.RequestResponse{}
	if err != nil {
		log.Fatal(err)
		response.Error = "No se pudo agregar referencia"
	} else {
		response.Succes = "Referencia agregada correctamente"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
func GetReference(w http.ResponseWriter, r *http.Request) {
	cursor, err := variables.ReferencesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var references []variables.ReferencesData
	for cursor.Next(context.TODO()) {
		i := variables.ReferencesData{}
		cursor.Decode(&i)
		i.Name = strings.ToUpper(i.Name)
		references = append(references, i)
	}
	JSONreferences, err := json.Marshal(references)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONreferences))
}
func DeleteReference(w http.ResponseWriter, r *http.Request) {
	response := variables.RequestResponse{}
	reference := variables.ReferencesData{}
	err := json.NewDecoder(r.Body).Decode(&reference)
	if err != nil {
		log.Fatal(err)
	}
	result, err := variables.ReferencesCollection.DeleteOne(context.TODO(), bson.D{{"_id", reference.Id}})
	if err != nil {
		log.Fatal(err)
		response.Error = "No se pudo eliminar proveedor"
	}
	if result.DeletedCount > 0 {
		response.Succes = "Proveedor Eliminado con Éxito"
	} else {
		response.Error = "No se encontró proveedor"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
func EditReference(w http.ResponseWriter, r *http.Request) {
	reference := variables.ReferencesData{}
	err := json.NewDecoder(r.Body).Decode(&reference)
	if err != nil {
		log.Fatal(err)
	}

	_, err = variables.ReferencesCollection.UpdateOne(context.TODO(), bson.D{{"_id", reference.Id}}, bson.D{{"$set", bson.D{{"name", reference.Name}, {"phone", reference.Phone}, {"completePrice", reference.CompletePrice}, {"reducedPrice", reference.ReducedPrice}, {"olderPrice", reference.OlderPrice}}}})
	response := variables.RequestResponse{}
	if err != nil {
		response.Error = "No se pudo editar referencia"
	} else {
		response.Succes = "Referencia editada correctamente"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))

}
