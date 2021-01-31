package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/byunjuneseok/hiupt/src/log"
)

func loginHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger(request.Body)

	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
}

func main()  {
	lambda.Start(loginHandler)
}