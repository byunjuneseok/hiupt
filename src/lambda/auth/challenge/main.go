package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/byunjuneseok/hiupt/src/auth"
	"github.com/byunjuneseok/hiupt/src/users"
)

func challengeHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// before processing request, get userId.
	var tmpUser users.User
	_ = json.Unmarshal([]byte(request.Body), &tmpUser)
	thisUserId := tmpUser.Id

	// validate userId.
	thisUser, err := users.Retrieve(thisUserId)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	// generate token.
	token, err := auth.CreateToken(thisUserId)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	result, err := auth.AuthChallengeQueueGetUrl(sess)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}
	queueUrl := result.QueueUrl

	err = auth.AuthChallengeQueueSendMessage(sess, queueUrl, thisUser, token)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: "Success.", StatusCode: 200}, nil
}


func main()  {
	lambda.Start(challengeHandler)
}