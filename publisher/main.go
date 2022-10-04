package main

import (
	"context"
	"log"
	"time"

	"github.com/go-zeromq/zmq4"
)

func main() {
	publisher := zmq4.NewPub(context.Background())

	defer publisher.Close()

	publisher.Listen("tcp://*:9092")

	for _ = range time.Tick(time.Second) {
		msg := zmq4.NewMsgFrom(
			[]byte("A"),
			[]byte(time.Now().Local().String()),
		)
		publisher.Send(msg)
		log.Println("send", "msg")
	}
}
