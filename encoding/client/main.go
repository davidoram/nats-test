package main

import (
	"log"
	"time"

	"github.com/davidoram/nats-test/encoding/data"
	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	for {
		a := data.Answer{}
		// Publish the message
		if err := ec.Request("hello", &data.Person{Firstname: "Dave", Surname: "Oram"}, &a, time.Second*1); err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %+v", a)
	}
}
