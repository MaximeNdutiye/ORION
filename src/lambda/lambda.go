package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
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
	sess := session.New()

	wab := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)
	uploader := s3manager.NewUploader(sess)

	_, dlErr := downloader.Download(wab, &s3.GetObjectInput{
		Bucket: aws.String("orion-image-bucket"),
		Key:    aws.String("image.jpg"),
	})

	imgBytes := bytes.NewReader(wab.Bytes())
	image, _, imgScalingErr := image.Decode(imgBytes)

	newImage := resize.Resize(160, 0, image, resize.Lanczos3)
	scaledImageBuff := bytes.NewBuffer(nil)
	jpgErr := jpeg.Encode(scaledImageBuff, newImage, nil)

	result, imgUploadErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("orion-image-bucket"),
		Key:    aws.String("image-scaled.jpg"),
		Body:   scaledImageBuff,
	})

	errors := [4]error{dlErr, imgScalingErr, jpgErr, imgUploadErr}

	for err := 0; err < 4; err++ {
		fmt.Println(errors[err])
	}

	fmt.Println("Finished image scale and upload")
}

func main() {
	lambda.Start(Handler)
}
