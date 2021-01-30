package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/byunjuneseok/hiupt/src/log"
)

func retrieveHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)  {
	log.Logger(request.Body)
	user, err := Retrieve(request.PathParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	jsonUser, _ := json.Marshal(user)
	stringUser := string(jsonUser) + "\n"

	return events.APIGatewayProxyResponse{Body: stringUser, StatusCode: 200}, err
}

func main()  {
	lambda.Start(retrieveHandler)
}