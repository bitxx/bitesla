package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:12345", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
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
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}
