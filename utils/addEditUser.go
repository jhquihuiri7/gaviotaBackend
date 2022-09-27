package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strings"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := variables.UsersData{}
	response := variables.RequestResponse{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = GenerateID()
	user.Password = EncryptPassword(user.Password)
	user.User = strings.ToUpper(user.Name[0:1] + user.LastName)
	//Identify if user already exists
	filter := bson.D{{"user", user.User}}
	err := variables.UsersCollection.FindOne(context.TODO(), filter).Err()
	if err == mongo.ErrNoDocuments {
		_, err := variables.UsersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			response.Error = "No se pudo crear usuario"
		} else {
			response.Succes = "Usuario creado correctamente"
		}
	} else {
		response.Error = "Usuario ya existente"
	}
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := variables.UsersData{}
	response := variables.RequestResponse{}
	json.NewDecoder(r.Body).Decode(&user)
	opts := options.FindOne().SetProjection(bson.D{{"name", 1}, {"lastName", 1}, {"user", 1}, {"rol", 1}})
	filter := bson.D{{"_id", user.Id}}
	err := variables.UsersCollection.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		response.Error = "No se encontró usuario"
		JSONresponse, _ := json.Marshal(response)
		fmt.Fprintln(w, string(JSONresponse))
	} else {
		JSONuser, _ := json.Marshal(user)
		fmt.Fprintln(w, string(JSONuser))
	}
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	opts := options.Find().SetProjection(bson.D{{"name", 1}, {"lastName", 1}, {"id", 1}, {"user", 1}, {"rol", 1}})
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
func EditUser(w http.ResponseWriter, r *http.Request) {
	user := variables.UsersData{}
	userOld := variables.UsersData{}
	json.NewDecoder(r.Body).Decode(&user)
	opts := options.FindOne().SetProjection(bson.D{{"name", 1}, {"lastName", 1}})
	filter := bson.D{{"_id", user.Id}}
	replace := bson.D{}

	if user.Name != "" || user.LastName != "" {
		if (user.Name != "" && user.LastName != "") == false {
			err := variables.UsersCollection.FindOne(context.TODO(), filter, opts).Decode(&userOld)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	userTemp := ""
	if user.Name != "" {
		replace = append(replace, bson.E{"name", user.Name})
		userTemp = strings.ToUpper(user.Name[0:1] + userOld.LastName)
	}
	if user.LastName != "" {
		replace = append(replace, bson.E{"lastName", user.LastName})
		if userTemp != "" {
			userTemp = strings.ToUpper(user.Name[0:1] + user.LastName)
		} else {
			userTemp = strings.ToUpper(userOld.Name[0:1] + user.LastName)
		}
	}

	if userTemp != "" {
		replace = append(replace, bson.E{"user", userTemp})
	}
	if user.Password != "" {
		replace = append(replace, bson.E{"password", EncryptPassword(user.Password)})
	}
	if user.Rol != "" {
		replace = append(replace, bson.E{"rol", user.Rol})
	}
	update := bson.D{{"$set", replace}}
	single := variables.UsersCollection.FindOneAndUpdate(context.TODO(), filter, update)
	response := variables.RequestResponse{}
	if single.Err() != nil {
		response.Error = "No se pudo actualizar usuario"
	} else {
		response.Succes = "Usuario actualizado correctamente"
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
