package variables

import (
	"crypto/rsa"
	"github.com/cloudinary/cloudinary-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection
var ReferencesCollection *mongo.Collection
var FrequentsCollection *mongo.Collection
var FrequentsNewCollection *mongo.Collection
var ReservesGaviotaCollection *mongo.Collection
var ReservesOtherCollection *mongo.Collection
var ReservesExternalCollection *mongo.Collection
var BinReservesCollection *mongo.Collection
var PaymentsSystemHistory *mongo.Collection
var PaymentsReferenceHistory *mongo.Collection

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

var Links []Link

var CloudinaryStorage *cloudinary.Cloudinary