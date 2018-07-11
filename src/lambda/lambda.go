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

type query struct {
	Image  string
	Width  string
	Height string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	queryStringParams := request.QueryStringParameters
	queryStringStruct := processQueryStringParams(queryStringParams)
	queryString, err := json.Marshal(queryStringStruct)
	outputJson := jsonResponse{string(queryString), 200}
	outputString, _ := json.Marshal(outputJson)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	logRequestInfo(request)
	return events.APIGatewayProxyResponse{
		Body:       string(outputString),
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

func processQueryStringParams(queryStringParams map[string]string) query {
	queryStringStruct := query{}

	for key, value := range queryStringParams {
		if key == "image" {
			queryStringStruct.Image = value
		} else if key == "width" {
			queryStringStruct.Width = value
		} else if key == "height" {
			queryStringStruct.Height = value
		}
	}

	return queryStringStruct
}

func main() {
	lambda.Start(Handler)
}
