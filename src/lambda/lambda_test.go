package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	orionRestApi, getRestApiErr := getRestApiFromAWS("orion-rest-api")
	restApiUrl := constructApiExecURL(*orionRestApi.Id, "orion", "us-east-1")

	if getRestApiErr != nil {
		fmt.Println(getRestApiErr)
		return
	}

	tests := []struct {
		Image  string
		Width  string
		expect int
		err    error
	}{
		{
			Image:  "image.jpg",
			Width:  "160",
			err:    nil,
			expect: 200,
		},
	}

	for _, test := range tests {

		url := restApiUrl + "/?imgpath=" + test.Image + "&width=" + test.Width
		fmt.Println(url)
		resp, err := http.Get(url)
		assert.Equal(t, err, nil)
		assert.Equal(t, resp.StatusCode, 200)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response body" + string(body))
		removeObjectFromS3("scaled/" + test.Image)
	}
}

func getRestApiFromAWS(apiName string) (*apigateway.RestApi, error) {
	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	gw := apigateway.New(sess)

	output, err := gw.GetRestApis(&apigateway.GetRestApisInput{
		Limit:    nil,
		Position: nil,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		for index := 0; index < len(output.Items); index++ {
			if *output.Items[index].Name == apiName {
				return output.Items[index], nil
			}
		}
	}

	return nil, errors.New("could not find apigateway rest api " + apiName)
}

func constructApiExecURL(apiId string, apiName string, region string) string {
	return "https://" + apiId + ".execute-api." + region + ".amazonaws.com/" + apiName
}

func removeObjectFromS3(objectPath string) {
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("orion-image-bucket"),
		Key:    aws.String(objectPath),
	}

	_, err := svc.DeleteObject(input)
	if err != nil {
		fmt.Println(err)
	}
}
