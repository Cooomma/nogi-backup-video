package helpers

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type messageHelper struct {
	service *sqs.SQS
	url     string
}

func NewMessageHelper(sess *session.Session, queueURL string) *messageHelper {
	service := sqs.New(sess)
	return &messageHelper{
		service: service,
		url:     queueURL,
	}
}

func (h messageHelper) SendMessage(content string) error {
	_, err := h.service.SendMessage(
		&sqs.SendMessageInput{
			MessageBody: &content,
			QueueUrl:    &h.url,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (h messageHelper) ReceiveMessage(maxNumberOfMessages int64) (*[]string, error) {
	output, err := h.service.ReceiveMessage(
		&sqs.ReceiveMessageInput{
			QueueUrl:            &h.url,
			MaxNumberOfMessages: &maxNumberOfMessages,
		},
	)
	if err != nil {
		return &[]string{}, err
	}
	messages := make([]string, len(output.Messages))
	for _, message := range output.Messages {
		messages = append(messages, message.GoString())
	}
	return &messages, nil
}
