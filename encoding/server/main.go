package main

import (
	"fmt"
	"log"
	"sync"

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

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Define the object
	type person struct {
		Firstname string
		Surname   string
	}

	type answer struct {
		Msg string
	}

	msgCnt := 0
	// Create a queue subscription on "updates" with queue name "workers"
	if _, err := ec.QueueSubscribe("hello", "hello-workers", func(subject, reply string, p *person) {
		log.Printf("Got %+v", p)

		a := answer{Msg: fmt.Sprintf("Hello %s %s", p.Firstname, p.Surname)}
		err = ec.Publish(reply, a)
		if err != nil {
			log.Fatal(err)
		}
		msgCnt += 1
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for messages to come in
	wg.Wait()
}
