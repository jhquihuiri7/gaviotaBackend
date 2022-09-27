package variables

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Login
type UserLogin struct {
	User     string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
}
type Claim struct {
	Credentials UserCredential
	jwt.StandardClaims
}
type ChangePasswordUser struct {
	Path string `json:"path"`
}
type RequestResponse struct {
	Error  string `json:"error"`
	Succes string `json:"succes"`
}
type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
	User  string `json:"user"`
	Rol   string `json:"rol"`
}

//Countries
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

//Users
type UsersData struct {
	Id       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	LastName string `bson:"lastName" json:"lastName"`
	User     string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
	Rol      string `bson:"rol" json:"rol"`
}

type UserCredential struct {
	User string `bson:"user" json:"user"`
	Rol  string `bson:"rol" json:"rol"`
}

//References
type ReferencesData struct {
	Id   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
}

//Reserves
type Reserve struct {
	Id string `bson:"_id" json:"id"`
	Reserve string 	`bson:"reserve" json:"reserve"`
	Passenger string `bson:"passenger" json:"passenger"`
	Referece string `bson:"referece" json:"referece"`
	User string `bson:"user" json:"user"`
	Age int `bson:"age" json:"age"`
	Date primitive.DateTime `bson:"date" json:"date"`
	Time string `bson:"time" json:"time"`
	Passport string `bson:"passport" json:"passport"`
	Country string `bson:"country" json:"country"`
	Price int `bson:"price" json:"price"`
	Ship string `bson:"ship" json:"ship"`
	Route string `bson:"route" json:"route"`
	IsConfirmed bool `bson:"isConfirmed" json:"isConfirmed"`
	IsPayed bool `bson:"isPayed" json:"isPayed"`
}



