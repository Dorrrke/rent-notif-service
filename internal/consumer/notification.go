package consumer

import (
	"encoding/json"
	"log"

	"github.com/Dorrrke/rent-notif-service/internal/models"
	"github.com/Dorrrke/rent-notif-service/internal/service"
	"github.com/nats-io/nats.go"
)

type Consumer struct {
	nc  *nats.Conn
	svc *service.NotificationService
}

func NewConsumer(nc *nats.Conn, svc *service.NotificationService) *Consumer {
	return &Consumer{
		nc:  nc,
		svc: svc,
	}
}

func (c *Consumer) Subscribe() error {
	_, err := c.nc.Subscribe(
		"user.registered",
		c.hadleUserRegistered,
	)
	if err != nil {
		return err
	}

	_, err = c.nc.Subscribe(
		"rent.started",
		c.handleRentStarted,
	)
	if err != nil {
		return err
	}

	_, err = c.nc.Subscribe(
		"rent.finished",
		c.handleRentFinished,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Consumer) hadleUserRegistered(msg *nats.Msg) {
	var event models.UserRegistered

	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("failed to unmarshal register event: %s", err.Error())
		return
	}

	if err := c.svc.SendWelcomeEmail(event.Email, event.Login); err != nil {
		log.Printf("failed to send welcome email: %s", err.Error())
		return
	}
}

func (c *Consumer) handleRentStarted(msg *nats.Msg) {
	var event models.RentStarted

	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Println(err)
		return
	}

	log.Printf("user id: %s", event.UserID)
	if err := c.svc.RentStarted(event.UserID, event.CarID, event.RentTime); err != nil {
		log.Println(err)
		return
	}
}

func (c *Consumer) handleRentFinished(msg *nats.Msg) {
	var event models.RentFinished

	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Println(err)
		return
	}

	if err := c.svc.RentFinished(event.UserID, event.CarID, event.Email, event.TotalRentTime); err != nil {
		log.Println(err)
		return
	}
}
