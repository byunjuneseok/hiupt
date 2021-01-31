package auth

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/byunjuneseok/hiupt/src/users"

	"os"
)

func AuthChallengeQueueGetUrl(sess *session.Session) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)

	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(os.Getenv("QUEUE_NAME_AUTH_CHALLENGE_EMAIL")),
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func AuthChallengeQueueSendMessage(sess *session.Session, queueUrl *string, user users.User, token string) error {
	svc := sqs.New(sess)
	var _, err = svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"userId": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(user.Id),
			},
			"userEmail": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(user.Email),
			},
			"token": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(token),
			},
		},
		MessageBody: aws.String("Login challenge email"),
		QueueUrl:    queueUrl,
	})
	if err != nil {
		return err
	}

	return nil
}

