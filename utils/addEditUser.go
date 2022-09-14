package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := variables.UsersData{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = GenerateID()
	user.Password = "123456"
	user.User = strings.ToUpper(user.Name[0:1] + user.LastName)
	_, err := variables.UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	opts := options.Find().SetProjection(bson.D{{"name", 1}, {"lastName", 1}, {"id", 1}})
	cursor, err := variables.UsersCollection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var users []variables.UsersData
	for cursor.Next(context.TODO()) {
		i := variables.UsersData{}
		cursor.Decode(&i)
		users = append(users, i)
	}
	JSONusers, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONusers))
}
func ChangePasswordUser(w http.ResponseWriter, r *http.Request) {
	user := variables.UsersData{}
	json.NewDecoder(r.Body).Decode(&user)
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", user.Id}}
	update := bson.D{{"$set", bson.D{{"password", EncryptPassword(user.Password)}}}}
	single := variables.UsersCollection.FindOneAndUpdate(context.TODO(), filter, update, opts)
	var response variables.RequestResponse
	if single.Err() != nil {
		response.Error = "No se pudo cambiar la contraseña"
	} else {
		response.Succes = "Contraseña cambiada con éxito"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	response := variables.RequestResponse{}
	user := variables.UsersData{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	result, err := variables.UsersCollection.DeleteOne(context.TODO(), bson.D{{"_id", user.Id}})
	if err != nil {
		log.Fatal(err)
		response.Error = "No se pudo eliminar usuario"
	}
	if result.DeletedCount > 0 {
		response.Succes = "Usuario Eliminado con Éxito"
	} else {
		response.Error = "No se encontró usuario"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
