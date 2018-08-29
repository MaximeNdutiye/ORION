gopath = /go
dir = $(shell pwd)

setupTravis:
	chmod +x setupTravis.sh && ./setupTravis.sh

setupAWS:
	chmod +x setupAWS.sh && ./setupAWS.sh

getDependencies:
	go get -t ./...

lambda: getDependencies
	cd src/lambda && GOOS=linux go build -o ../../build/orion

zip: lambda
	cd build && zip lambda.zip orion

test:
	cd src/lambda && go test -v -race ./...

container:
	docker build -t orion .

runcontainer:
	docker run -it -v $(dir):$(gopath)/ --rm orion

deployLambda:
	cd src/aws && ../../terraform init && ../../terraform apply -auto-approve

destroyLambda:
	cd src/aws && ../../terraform destroy -auto-approve

deploy: zip
	cd src/aws && terraform init && terraform apply

destroy:
	cd src/aws && terraform destroy
