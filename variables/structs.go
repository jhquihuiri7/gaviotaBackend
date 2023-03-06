package variables

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Login
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
type RequestResponseNumber struct {
	Error  int `json:"error"`
	Succes int `json:"succes"`
}
type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
	User  string `json:"user"`
	Rol   string `json:"rol"`
}

// Countries
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Users
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

// References
type ReferencesData struct {
	Id            string `bson:"_id" json:"id"`
	Name          string `bson:"name" json:"name"`
	Phone         string `bson:"phone" json:"phone"`
	CompletePrice int    `bson:"completePrice" json:"completePrice"`
	ReducedPrice  int    `bson:"reducedPrice" json:"reducedPrice"`
	OlderPrice    int    `bson:"olderPrice" json:"olderPrice"`
}

// Reserves
type Reserve struct {
	Id             string             `bson:"_id" json:"id"`
	ReserveNumber  string             `bson:"reserve" json:"reserve"`
	Passenger      string             `bson:"passenger" json:"passenger"`
	Reference      string             `bson:"reference" json:"reference"`
	User           string             `bson:"user" json:"user"`
	Age            int                `bson:"age" json:"age"`
	Date           primitive.DateTime `bson:"date" json:"date"`
	Time           string             `bson:"time" json:"time"`
	Passport       string             `bson:"passport" json:"passport"`
	Country        string             `bson:"country" json:"country"`
	Price          int                `bson:"price" json:"price"`
	Ship           string             `bson:"ship" json:"ship"`
	Route          string             `bson:"route" json:"route"`
	IsConfirmed    bool               `bson:"isConfirmed" json:"isConfirmed"`
	IsPayed        bool               `bson:"isPayed" json:"isPayed"`
	IsInvoiced     bool               `bson:"isInvoiced" json:"isInvoiced"`
	Birthday       primitive.DateTime `bson:"birthday" json:"birthday"`
	Comment        string             `bson:"comment" json:"comment"`
	Notes          string             `bson:"notes" json:"notes"`
	Status         string             `bson:"status" json:"status"`
	Phone          string             `bson:"phone" json:"phone"`
	PaymentMethod  string             `bson:"paymentMethod" json:"paymentMethod"`
	PaymentDate    primitive.DateTime `bson:"paymentDate" json:"paymentDate"`
	RegisteredDate primitive.DateTime `bson:"registeredDate" json:"registeredDate"`
}

type MultiplyReserve struct {
	Reserve
	Number int `json:"number"`
}

type EditedReserve struct {
	Reserve
	Collection string `json:"collection"`
}
type Ids struct {
	Id string `json:"id"`
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

// Report Scarlett
type ReportSalesFilter struct {
	InitDate   primitive.DateTime `json:"initDate"`
	FinalDate  primitive.DateTime `json:"finalDate"`
	Collection string             `json:"collection"`
}
type ReportSalesData struct {
	Reference string    `json:"reference" bson:"reference"`
	Id        string    `bson:"_id" json:"id"`
	Total     int       `json:"total"`
	Payed     int       `json:"payed"`
	Reserves  []Reserve `json:"reserves"`
}
type RefereceSalesState struct {
	Id          string               `bson:"_id" json:"id"`
	Name        string               `bson:"name" json:"name"`
	Balance     int                  `json:"balance"`
	Pending     int                  `json:"pending"`
	MonthlyData []RefereceSalesMonth `json:"monthlyData"`
}
type RefereceSalesMonth struct {
	Date          primitive.DateTime `json:"date"`
	TotalReserves int                `json:"totalReserves"`
	PayedReserves int                `json:"payedReserves"`
}

type AdvanceDataRequest struct {
	ReferenceName string `json:"referenceName"`
	User          string `bson:"user" json:"user"`
	Advance       int    `bson:"advance" json:"advance"`
	PaymentMethod string `bson:"paymentMethod" json:"paymentMethod"`
}
type AdvanceData struct {
	User          string             `bson:"user" json:"user"`
	Advance       int                `bson:"advance" json:"advance"`
	Pending       int                `bson:"pending" json:"pending"`
	Balance       int                `bson:"balance" json:"balance"`
	Total         int                `bson:"total" json:"total"`
	PaymentMethod string             `bson:"paymentMethod" json:"paymentMethod"`
	Date          primitive.DateTime `bson:"date" json:"date"`
	PayedIds      []string           `bson:"payedIds" json:"payedIds"`
}

type ReferencePaymentHistory struct {
	Id        string `bson:"_id" json:"id"`
	Reference string `bson:"reference" json:"reference"`
	History   []AdvanceData
}

type Link struct {
	LinkId       string
	CreationDate time.Time
}
type TicketRoute struct {
	Date  primitive.DateTime `bson:"date" json:"date"`
	Route string             `bson:"route" json:"route"`
	Time  string             `bson:"time" json:"time"`
	Ship  string             `bson:"ship" json:"ship"`
}
type TicketPassenger struct {
	Passenger string `bson:"passenger" json:"passenger"`
	Country   string `bson:"country" json:"country"`
}
type Ticket struct {
	Routes     [][]string
	Passengers [][]string
	RouteId    []string
	Total      int
}

type DailyUserSales struct {
	User                  string `json:"user"`
	Total                 int    `json:"total"`
	NotPayed              int    `json:"notPayed"`
	Cash                  int    `json:"cash"`
	CreditDebitCard       int    `json:"creditDebitCard"`
	WireTransferPichincha int    `json:"wireTransferPichincha"`
	WireTransferPacifico  int    `json:"wireTransferPacifico"`
	DepositPichincha      int    `json:"depositPichincha"`
	DepositPacifico       int    `json:"depositPacifico"`
}
type DailyIncomes struct {
	Total    int       `json:"total"`
	Method   string    `json:"method"`
	Reserves []Reserve `json:"reserves"`
}

// DEV
type FrequentsOld struct {
	Name     string `bson:"referencia"`
	Status   string `bson:"status"`
	Country  string `bson:"nacionalidad" json:"country"`
	Passport string `bson:"cedula" json:"passport"`
	Age      int    `bson:"edad"`
}
type Frequent struct {
	Id       string             `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Status   string             `bson:"status" json:"status"`
	Country  string             `bson:"country" json:"country"`
	Passport string             `bson:"passport" json:"passport"`
	Birthday primitive.DateTime `bson:"birthday" json:"birthday"`
	Age      int                `json:"age"`
	Phone    string             `bson:"phone" json:"phone"`
}
