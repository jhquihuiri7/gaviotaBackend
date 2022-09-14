package authentication

import (
	"encoding/json"
	"fmt"
	"gaviotaBackend/variables"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"log"
	"net/http"
)

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	response := variables.LoginResponse{}
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return variables.PublicKey, nil
	}, request.WithClaims(&variables.Claim{}))
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				response.Error = "Id Expired"
			case jwt.ValidationErrorSignatureInvalid:
				response.Error = "No valid token"
			}
		default:
			response.Error = "No valid token"
		}
	}
	if token.Valid {
		payload := DecodeJWT(token.Raw)
		response.User = payload.User
		response.Rol = payload.Rol
		response.Token, response.Error = GenerateJWT(payload)

	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(JSONresponse))
}
