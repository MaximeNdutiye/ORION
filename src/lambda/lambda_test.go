package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	queryStringParamsMap := make(map[string]string)
	tests := []struct {
		request events.APIGatewayProxyRequest
		Image   string
		Width   string
		Height  string
		expect  int
		err     error
	}{
		{
			// Send a request to API gateway with some QueryStringParameters
			request: events.APIGatewayProxyRequest{QueryStringParameters: queryStringParamsMap},
			Image:   "hello.jpg",
			Width:   "100",
			Height:  "200",
			expect:  200,
			err:     nil,
		},
	}

	for _, test := range tests {
		queryStringParamsMap["image"] = test.Image
		queryStringParamsMap["width"] = test.Width
		queryStringParamsMap["height"] = test.Height

		response, err := Handler(test.request)
		fmt.Println("Response: ", response)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.StatusCode)
	}
}
