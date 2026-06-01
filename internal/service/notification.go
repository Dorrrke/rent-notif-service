package service

type EmailSender interface {
	SendMail(to, subject, body string) error
}

type PushSender interface {
	SendPush(userID, title, body string) error
}

type NotificationService struct {
	emailSender EmailSender
	pushSender  PushSender
}

func NewNotificationService(emailSender EmailSender, pushSender PushSender) *NotificationService {
	return &NotificationService{
		emailSender: emailSender,
		pushSender:  pushSender,
	}
}

func (s *NotificationService) SendWelcomeEmail(email, login string) error {
	// log.Printf("[EMAIL] Welcome to the service, %s (%s)", email, login)

	return s.emailSender.SendMail(email, "Welcome to the service", "Welcome to the service, "+login)
}

func (s *NotificationService) RentStarted(
	userID,
	carID,
	rentTime string,
) error {
	return s.pushSender.SendPush(userID, "Rent started", "You have started a rent of "+carID+" for "+rentTime)
}

func (s *NotificationService) RentFinished(
	userID,
	carID,
	email,
	totalTime string,
) error {
	err := s.emailSender.SendMail(email, "Rent finished", "You have finished a rent of "+carID+" for "+totalTime)
	if err != nil {
		return err
	}

	return s.pushSender.SendPush(userID, "Rent finished", "You have finished a rent of "+carID+" for "+totalTime)
}
