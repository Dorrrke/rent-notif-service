package main

import (
	"fmt"
	"log"

	"github.com/Dorrrke/rent-notif-service/internal/consumer"
	"github.com/Dorrrke/rent-notif-service/internal/service"
	pushMock "github.com/Dorrrke/rent-notif-service/internal/transport/pushapi/mock"
	"github.com/Dorrrke/rent-notif-service/internal/transport/smtp/mock"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	push := pushMock.NewPushAPIMock()
	smtp := mock.NewEmailMockSender()

	notifSvc := service.NewNotificationService(smtp, push)

	consumer := consumer.NewConsumer(nc, notifSvc)

	if err := consumer.Subscribe(); err != nil {
		log.Fatal(err)
	}

	log.Println("service started")
	fmt.Println("---------------------------------------------------------------------")

	select {}
}
