package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type jsonResponse struct {
	Parameters string
	Status     int
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	output := fmt.Sprintf("parameters", request.QueryStringParameters)
	outputResponse := jsonResponse{output, 200}
	resp, err := json.Marshal(successObj)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	logRequestInfo(request)
	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func logRequestInfo(request events.APIGatewayProxyRequest) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	fmt.Println("Query String Parameters: ")
	for key, value := range request.QueryStringParameters {
		fmt.Printf("	%s: %s\n", key, value)
	}
}

func main() {
	lambda.Start(Handler)
}
