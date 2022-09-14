package variables

import "github.com/golang-jwt/jwt"

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
	Succes string `json:"token"`
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
