package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type SendHandler struct{}

func NewSendHandler() *SendHandler {
	return &SendHandler{}
}

func (sh *SendHandler) HandleRequest(ctx context.Context, event map[string]interface{}) error {
	log.Println("start handling event.", event)
	return nil
}

func main() {
	handler := NewSendHandler()
	lambda.Start(handler.HandleRequest)
}
