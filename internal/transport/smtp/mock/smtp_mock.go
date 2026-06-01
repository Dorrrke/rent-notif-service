package mock

import (
	"fmt"
	"log"
)

type EmailMockSender struct{}

func NewEmailMockSender() *EmailMockSender {
	return &EmailMockSender{}
}

func (s *EmailMockSender) SendMail(to, subject, body string) error {
	log.Printf("[EMAIL NOTIFICATION]\nFrom: gorent.support\nTo: %s\nSubject:%s\nBody:%s", to, subject, body)
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println()
	return nil
}
