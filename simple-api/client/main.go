package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create a unique subject name for replies.
	uniqueReplyTo := nats.NewInbox()

	// Listen for a single response
	sub, err := nc.SubscribeSync(uniqueReplyTo)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Send the request.
		// If processing is synchronous, use Request() which returns the response message.
		if err := nc.PublishRequest("names", uniqueReplyTo, nil); err != nil {
			log.Fatal(err)
		}

		// Read the reply
		msg, err := sub.NextMsg(time.Second)
		if err != nil {
			log.Fatal(err)
		}

		// Use the response
		log.Printf("Reply: %s", msg.Data)

	}
}
