package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

type jsonResponse struct {
	Parameters string
	Status     int
}

type query struct {
	Image  string
	Width  string
	Height string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryStringParameters, err := json.Marshal(request.QueryStringParameters)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	logRequestInfo(request)
	return events.APIGatewayProxyResponse{
		Body:       string(queryStringParameters),
		StatusCode: 200,
	}, nil
}

func logRequestInfo(request events.APIGatewayProxyRequest) {
	fmt.Printf("Request Id %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	fmt.Println("QueryStringParameters: ")
	for key, value := range request.QueryStringParameters {
		fmt.Printf("	%s: %s\n", key, value)
	}
}

func main() {
	lambda.Start(Handler)
}
