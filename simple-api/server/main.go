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

	wg := sync.WaitGroup{}
	wg.Add(1)

	msgCnt := 0
	// Create a queue subscription on "updates" with queue name "workers"
	if _, err := nc.QueueSubscribe("names", "name-workers", func(m *nats.Msg) {
		m.Respond([]byte(fmt.Sprintf("hello %d", msgCnt)))
		msgCnt += 1
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for messages to come in
	wg.Wait()
}
