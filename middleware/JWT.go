package middleware

import (
	"gaviotaBackend/variables"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"net/http"
	"strings"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/api/login" && r.URL.String() != "/api/validateToken" && r.URL.String() != "/api/addReservesExternal" && r.URL.String() != "/admin/addPrices" && strings.Contains(r.URL.String(), "/api/dailyReport") != true && strings.Contains(r.URL.String(), "/response") != true && r.URL.String() != "/api/generateLink" && strings.Contains(r.URL.String(), "/api/validateLink") != true && strings.Contains(r.URL.String(), "/api/generateTicket") != true && strings.Contains(r.URL.String(), "/api/GaviotaTicketFerry") != true && strings.Contains(r.URL.String(), "/api/printTicket") != true && strings.Contains(r.URL.String(), "/api/ShareOtherReserves") != true && r.URL.String() != "/order" {
			token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
				return variables.PublicKey, nil
			}, request.WithClaims(&variables.Claim{}))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if token.Valid {
				//w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusOK)
		}
		// Next
		next.ServeHTTP(w, r)
		return
	})
}
