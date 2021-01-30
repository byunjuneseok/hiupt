package users

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/byunjuneseok/hiupt/api"
)

func retrieveHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)  {
	api.Logger(request.Body)
	//_, err := Retrieve(request.Body)

	return events.APIGatewayProxyResponse{}, nil
}

func main()  {
	lambda.Start(retrieveHandler)
}