package pkg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type IEmailService interface {
	SendEmail(from, to, subject, body string) error
}

type emailSender struct{}

func NewEmailSender() IEmailService {
	return &emailSender{}
}

func (es *emailSender) SendEmail(from, to, subject, body string) error {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		return err
	}

	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: &from,
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &body,
					},
				},
				Subject: &types.Content{
					Data: &subject,
				},
			},
		},
	}

	_, err = client.SendEmail(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
