package mock

import (
	"fmt"
	"log"
)

type PushAPIMock struct{}

func NewPushAPIMock() *PushAPIMock {
	return &PushAPIMock{}
}

func (p *PushAPIMock) SendPush(userID, title, body string) error {
	log.Printf("[SEND PUSH NOTIFICATION]\nTitle: %s\nBody:%s\nTo:%s", title, body, userID)
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println()
	return nil
}
