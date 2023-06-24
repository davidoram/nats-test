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

	ec, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type person struct {
		Firstname string
		Surname   string
	}

	type answer struct {
		Msg string
	}
	a := answer{}
	// Publish the message
	if err := ec.Request("hello", &person{Firstname: "Dave", Surname: "Oram"}, &a, time.Second*1); err != nil {
		log.Fatal(err)
	}
	log.Printf("Reply: %+v", a)
}
