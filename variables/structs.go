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
	Id            string `bson:"_id" json:"id"`
	ReserveNumber string `bson:"reserve" json:"reserve"`
	Passenger     string `bson:"passenger" json:"passenger"`
	Referece    string             `bson:"referece" json:"referece"`
	User        string             `bson:"user" json:"user"`
	Age         int                `bson:"age" json:"age"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	Time        string             `bson:"time" json:"time"`
	Passport    string             `bson:"passport" json:"passport"`
	Country     string             `bson:"country" json:"country"`
	Price       int                `bson:"price" json:"price"`
	Ship        string             `bson:"ship" json:"ship"`
	Route       string             `bson:"route" json:"route"`
	IsConfirmed bool               `bson:"isConfirmed" json:"isConfirmed"`
	IsPayed     bool               `bson:"isPayed" json:"isPayed"`
	IsBlocked bool `bson:"isBlocked" json:"isBlocked"`
}

type MultiplyReserve struct {
	Reserve
	Number int `json:"number"`
}

type DailyReportRequest struct {
	Time string             `bson:"time" json:"time"`
	Date primitive.DateTime `bson:"date" json:"date"`
}

type PaymentInfo struct {
	Id                  string `bson:"_id" json:"id"`
	Email               string `bson:"email" json:"email"`
	CardType            string `bson:"cardType" json:"cardType"`
	Bin                 string `bson:"bin" json:"bin"`
	LastDigits          string `bson:"lastDigits" json:"lastDigits"`
	DeferredCode        string `bson:"deferredCode" json:"deferredCode"`
	Deferred            bool   `bson:"deferred" json:"deferred"`
	CardBrandCode       string `bson:"cardBrandCode" json:"cardBrandCode"`
	CardBrand           string `bson:"cardBrand" json:"cardBrand"`
	Amount              int64  `bson:"amount" json:"amount"`
	ClientTransactionID string `bson:"clientTransactionId" json:"clientTransactionId"`
	PhoneNumber         string `bson:"phoneNumber" json:"phoneNumber"`
	TransactionStatus   string `bson:"transactionStatus" json:"transactionStatus"`
	AuthorizationCode   string `bson:"authorizationCode" json:"authorizationCode"`
	TransactionID       int64  `bson:"transactionId" json:"transactionId"`
	Document            string `bson:"document" json:"document"`
	Currency            string `bson:"currency" json:"currency"`
	StoreName           string `bson:"storeName" json:"storeName"`
	Date                string `bson:"date" json:"date"`
	RegionISO           string `bson:"regionIso" json:"regionIso"`
	TransactionType     string `bson:"transactionType" json:"transactionType"`
	Reference           string `bson:"reference" json:"reference"`
}
