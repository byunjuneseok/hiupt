package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/byunjuneseok/hiupt/src/log"
	"github.com/byunjuneseok/hiupt/src/users"
)

func userCreateHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger(request.Body)

	_, err := users.Create(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Error",
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body: "Success",
		StatusCode: 200,
	}, nil
}

func main()  {
	lambda.Start(userCreateHandler)
}