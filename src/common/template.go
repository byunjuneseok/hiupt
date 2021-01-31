package common

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
}

func main()  {
	lambda.Start(Handler)
}