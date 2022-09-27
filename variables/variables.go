package variables

import (
	"crypto/rsa"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection
var ReferencesCollection *mongo.Collection
var FrequentsCollection *mongo.Collection
var ReservesGaviotaCollection *mongo.Collection
var ReservesOtherCollection *mongo.Collection
var BinReservesCollection *mongo.Collection

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey





