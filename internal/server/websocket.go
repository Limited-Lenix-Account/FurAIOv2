package server

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func ConnectToSocket() (*websocket.Conn, error) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "SERVER_URL", Path: "/skus"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		// log.Fatal("dial:", err)
		return nil, fmt.Errorf("error dialing websocket: %w", err)
	}

	return c, nil

}

func ListenToSocket(c *websocket.Conn, messageChan chan<- []byte) {
	log.Println("Listening to socket")

	defer func() {
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("Error closing connection:", err)
		}
		c.Close()
	}()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error recieving message:", err)
			return
		}
		// log.Printf("Received message: %s", message)
		messageChan <- message
	}
}

func ListenToSocketWithReconnect(url string, messageChan chan []byte) {
	for {
		err := listenToSocket(url, messageChan)
		if err != nil {
			log.Printf("WebSocket error: %v", err)
			time.Sleep(5 * time.Second) // Wait before reconnecting
		}
	}
}

func listenToSocket(url string, messageChan chan []byte) error {
	fmt.Println("listening to socket")
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			log.Println("Interrupt received, shutting down...")
			return nil
		default:
			fmt.Println("reading message")
			_, message, err := conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("read error: %w", err)
			}
			messageChan <- message
		}
	}
}
