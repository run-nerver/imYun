package main

import (
	"context"
	"fmt"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
	"log"
	"time"
)

func main()  {
	ctx := context.TODO()
	events := make(neffos.Namespaces)
	events.On("default", "Notify", func(c *neffos.NSConn, msg neffos.Message) error {
		log.Printf("Server says: %s\n", string(msg.Body))
		neffos.Reply([]byte("OK"))
		c.Emit("Notify", []byte("Pong!"))
		return nil
	})

	// Connect to the server.
	client, err := neffos.Dial(ctx,
		gorilla.DefaultDialer,
		"ws://localhost:5000/v1/webs/websocket",
		events)
	if err != nil {
		panic(err)
	}
	// Connect to a namespace.
	c, err := client.Connect(ctx, "default")
	if err != nil {
		panic(err)
	}
	Tokens := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTEwMzk1NjYsImlhdCI6MTYxMDk1MzE2NiwiaWQiOjIsImlzcyI6IklyaXMiLCJ1c2VybmFtZSI6ImFkbWluMSJ9.tGuRZeF9oBLsAeO4IyQQQmaDMqGypk0UhPQxqVNio1c")
	response, err := c.Ask(ctx, "Authorization",Tokens)

	if err != nil {
		if neffos.IsCloseError(err) {
			fmt.Printf("error received: %v\n", err)
			c.Disconnect(ctx)
		}
	}
	fmt.Println(string(response.Body))
	for {
		time.Sleep(1*time.Minute)
		c.Ask(ctx, "Authorization", []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTA5NjU4NjUsImlhdCI6MTYxMDg3OTQ2NSwiaWQiOjEsImlzcyI6IklyaXMiLCJ1c2VybmFtZSI6ImFkbWluIn0.UkiRRr3CAB9jB15RT7571YocS05sQ_-SMezUU4vDMZA"))

		fmt.Println(string(response.Body))
	}
}