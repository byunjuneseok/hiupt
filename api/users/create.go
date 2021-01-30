package users

import (
	"github.com/byunjuneseok/hiupt/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	api.Logger(request.Body)

	_, err := Create(request.Body)

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
	lambda.Start(Handler)
}