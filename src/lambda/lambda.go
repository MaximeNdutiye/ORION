package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
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

	getObjectFromBucket()

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

func getObjectFromBucket() {
	svc := s3.New(session.New())
	input := &s3.GetObjectInput{
		Bucket: aws.String("orion-config-bucket"),
		Key:    aws.String("config.json"),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func main() {
	lambda.Start(Handler)
}
