
lambda:
	cd ../src/lambda && GOOS=linux go build -o ../../build/orion

build: lambda
	cd ../build && zip orion.zip orion