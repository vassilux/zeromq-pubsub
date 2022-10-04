package main

import (
	"context"
	"log"

	"github.com/go-zeromq/zmq4"
)

func main() {
	subscriber := zmq4.NewSub(context.Background())

	defer subscriber.Close()

	err := subscriber.Dial("tcp://127.0.0.1:9092")

	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	err = subscriber.SetOption(zmq4.OptionSubscribe, "A")
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}
	for {
		// Read envelope
		msg, err := subscriber.Recv()
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}
		log.Printf("[%s] %s\n", msg.Frames[0], msg.Frames[1])
	}

}
