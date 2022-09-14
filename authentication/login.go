package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userData variables.UserLogin
	var userPassword variables.UsersData
	var response variables.LoginResponse
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		log.Fatal(err)
	}
	opts := options.FindOne().SetProjection(bson.D{{"rol", 1}, {"password", 1}})
	err = variables.UsersCollection.FindOne(context.TODO(), bson.D{{"user", strings.ToUpper(userData.User)}}, opts).Decode(&userPassword)
	if err != nil {
		response.Error = "Usuario no registrado"
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(userPassword.Password), []byte(userData.Password))
		if err == nil {
			credentials := variables.UserCredential{
				User: userData.User,
				Rol:  userPassword.Rol,
			}
			response.Token, response.Error = GenerateJWT(credentials)
			response.User = userData.User
			response.Rol = userPassword.Rol
		} else {
			response.Error = "Contraseña incorrecta"
		}
	}
	JSONresponse, _ := json.Marshal(response)
	fmt.Fprintln(w, string(JSONresponse))
}
