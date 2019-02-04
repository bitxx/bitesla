package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:2303", "http service address")

func main() {
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ping"}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	}
}
