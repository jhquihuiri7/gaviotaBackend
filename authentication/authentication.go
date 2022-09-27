package authentication

import (
	"fmt"
	"gaviotaBackend/variables"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"time"
)

func InitKeys() {
	//Read Private and Public Key from directory
	privateBytes, _ := ioutil.ReadFile("files/private.pem")
	publicBytes, _ := ioutil.ReadFile("files/public.pem")
	//Parse Private and Public Key from []Byte
	variables.PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	variables.PublicKey, _ = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
}
func GenerateJWT(credentials variables.UserCredential) (string, string) {
	claim := variables.Claim{
		Credentials: credentials,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			Issuer:    "Logiciel Applab",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	result, err := token.SignedString(variables.PrivateKey)
	if err != nil {
		return "", err.Error()
	}
	return result, ""
}
func DecodeJWT(tokenString string) variables.UserCredential {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	payload := claims["Credentials"]
	payloadMap := payload.(map[string]interface{})
	return variables.UserCredential{User: fmt.Sprintf("%v", payloadMap["user"]), Rol: fmt.Sprintf("%v", payloadMap["rol"])}
}
