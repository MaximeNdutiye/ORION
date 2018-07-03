lambda:
	cd src/lambda && GOOS=linux go build -o ../../build/orion

zip: lambda
	cd build && zip lambda.zip orion

test:
	cd src/lambda && go test -v -race ./...

getDependencies:
	cd src/lambda && go get ./...

container:
	docker build -t orion .
