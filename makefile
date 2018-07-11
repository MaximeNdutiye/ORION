gopath = /go
dir = $(shell pwd)

setupAWS:
	chmod +x setupAWS.sh && ./setupAWS.sh

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

runcontainer:
	docker run -it -v $(dir):$(gopath)/ --rm orion

deployLambda:
	cd src/aws terraform apply -auto-approve

destroyLambda:
	cd src/aws && terraform destroy -auto-approve

setupTravis:
	chmod +x setupTravis.sh && ./setupTravis.sh