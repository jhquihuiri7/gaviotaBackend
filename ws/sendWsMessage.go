package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Dialer websocket.Dialer
var Connection *websocket.Conn

func WsConnection() *websocket.Conn {
	conn, _, err := Dialer.Dial("wss://websocket-microservice.herokuapp.com/api/wsGaviota", http.Request{}.Header)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}
func SendSocketMessage(reserves []interface{}) {
	data, _ := json.Marshal(reserves)
	fmt.Println(len(data))
	err := Connection.WriteJSON(reserves)
	if err != nil {
		fmt.Println(err)
	}
}
