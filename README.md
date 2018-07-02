# ORION
Serverless image scaling service built to run in AWS.

https://travis-ci.com/MaximeNdutiye/ORION\
[![Build Status](https://travis-ci.com/MaximeNdutiye/ORION.svg?token=jz2ngWM3JpnFiWRYz9Ru&branch=master)]

### Building the Docker images
Building the docker container will create and configure a container with everything needed for development.
The docker container is not needed, but it is useful since the environement is preconfigured.
`Note`: If the container isn't being used then the correct version of terraform and other tools must be installed by the user

`docker build -t orion .`

### Set up AWS-CLI
AWS credentials must be setup using aws-cli inside the container to be able to deploy to aws

`aws setup` will prompt for the credentials needed

## Development

### Running the container

`docker run -it --rm orion`

### Mouting directories
You can mount your own dirrectories to the container for development purposes

`docker run -it --rm -v /path/to/volume orion`

### Building the lamda
The lambda function must be package so it can be deployed by terraform

### Run Tests

`make test`

## FUTURE
Things that we want to get done

### Built with :purple_heart: By:
[Maxime](https://github.com/MaximeNdutiye)

### Links
[Deploying lambda function with terraform](https://www.terraform.io/docs/providers/aws/guides/serverless-with-aws-lambda-and-api-gateway.html)
[Lambda function handler in go](https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html)