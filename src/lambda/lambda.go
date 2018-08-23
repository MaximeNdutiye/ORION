package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
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
	sess := session.New()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	bucketName := "orion-image-bucket"
	imagePath := request.QueryStringParameters["imgpath"]
	desiredWidth, _ := strconv.ParseUint(request.QueryStringParameters["width"], 10, 32)
	image := getObjectFromS3(bucketName, imagePath, sess)
	scaledImgBuff := scaleImage(image, uint(desiredWidth))
	updloadToS3(scaledImgBuff, sess)

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

func getObjectFromS3(bucketName string, objectPath string, sess *Session) image.Image {
	wab := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)

	_, dlErr := downloader.Download(wab, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectPath),
	})

	if dlErr != nil {
		fmt.Println(dlErr)
		return
	}

	imgBytes := bytes.NewReader(wab.Bytes())
	img, _, imgDecodeErr := image.Decode(imgBytes)

	if imgDecodeErr != nil {
		fmt.Println(imgDecodeErr)
		return
	}

	return img
}

func scaleImage(img image.Image, desiredWidth uint) *Buffer {
	newImage := resize.Resize(desiredWidth, 0, img, resize.Lanczos3)
	scaledImageBuff := bytes.NewBuffer(nil)
	jpgErr := jpeg.Encode(scaledImageBuff, newImage, nil)

	if jpgErr != nil {
		fmt.Println(jpgErr)
		return
	}

	return scaledImageBuff
}

func updloadToS3(imageBuffer *Buffer, sess *Session) {
	uploader := s3manager.NewUploader(sess)
	_, imgUploadErr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("scaled/" + objectPath),
		Body:   imageBuffer,
	})

	if imgUploadErr != nil {
		fmt.Println(imgUploadErr)
	}
}

func main() {
	lambda.Start(Handler)
}
