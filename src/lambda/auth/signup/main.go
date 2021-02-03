package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/byunjuneseok/hiupt/src/utils"
	"github.com/byunjuneseok/hiupt/src/log"
	"github.com/byunjuneseok/hiupt/src/users"
)

func signUpHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger(request.Body)

	thisUser, err := users.Create(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(),
			StatusCode: 500,
		}, nil
	}

	str, err := utils.ObjectToJsonString(thisUser)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body: str,
		StatusCode: 200,
	}, nil
}

func main()  {
	lambda.Start(signUpHandler)
}
