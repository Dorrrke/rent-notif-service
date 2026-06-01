package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	if err := nc.Publish("rent.started", []byte(`{"user_id": "123", "carID": "123", "rentTime": "123"}`)); err != nil {
		log.Fatal(err)
	}

	if err := nc.Publish("rent.finished", []byte(`{"user_id": "123", "carID": "123", "email": "123", "rentTime": "123"}`)); err != nil {
		log.Fatal(err)
	}

	if err := nc.Publish("user.registered", []byte(`{"user_id": "123", "email": "123", "login": "123"}`)); err != nil {
		log.Fatal(err)
	}

}
