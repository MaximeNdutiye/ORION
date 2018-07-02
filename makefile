lambda:
	cd src/lambda && GOOS=linux go build -o ../../build/orion

zip: lambda
	cd build && zip lambda.zip orion

install:
	cd src/lambda && go install

test:
	cd tests && go test -v -race ./...