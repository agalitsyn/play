package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Message struct {
	Name    int    `json:"name"`
	Message string `json:"message"`
}

// defaults https://socket.io/docs/rooms-and-namespaces/
const room = "messages"
const namespace = "/"

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Printf("client %s connected", c.Id())

		c.Join(namespace)

		tick := time.Tick(time.Second * 1)
		for now := range tick {
			c.Emit(room, Message{1, now.String()})
		}
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Printf("client %s disconnected", c.Id())
	})

	server.On("send", func(c *gosocketio.Channel, msg Message) string {
		c.BroadcastTo(namespace, room, msg)
		log.Printf("msg from %d: %s", msg.Name, msg.Message)
		return "OK"
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
