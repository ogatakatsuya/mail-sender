package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"mail-sender/pkg"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SendHandler struct {
	es pkg.IEmailService
}

type Event struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

var (
	from = os.Getenv("FROM_EMAIL")
	to   = os.Getenv("TO_EMAIL")
)

func NewSendHandler(es pkg.IEmailService) *SendHandler {
	return &SendHandler{
		es: es,
	}
}

func (sh *SendHandler) HandleRequest(ctx context.Context, event events.SQSEvent) error {
	log.Println("Handling event:", event)

	for _, message := range event.Records {
		var eventData Event
		if err := json.Unmarshal([]byte(message.Body), &eventData); err != nil {
			return errors.New("failed to unmarshal event data: " + err.Error())
		}
		err := sh.es.SendEmail(
			from,
			to,
			eventData.Subject,
			eventData.Body,
		)
		if err != nil {
			return errors.New("failed to send email: " + err.Error())
		}
	}

	log.Println("Finished handling event.", event)
	return nil
}

func main() {
	emailService := pkg.NewEmailSender()
	handler := NewSendHandler(emailService)
	lambda.Start(handler.HandleRequest)
}
