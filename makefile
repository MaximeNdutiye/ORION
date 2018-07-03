lambda:
	cd src/lambda && GOOS=linux go build -o ../../build/orion

zip: lambda
	cd build && zip lambda.zip orion

test:
	cd src/lambda && go test -v -race ./...

getDependencies:
	go get "github.com/aws/aws-lambda-go/events" "github.com/stretchr/testify/assert"

container:
	docker build -t orion .
