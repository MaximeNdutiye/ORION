# ORION
Serverless image scaling service built with `AWS Lambda` and `APIGateway`.

[![Build Status](https://travis-ci.com/MaximeNdutiye/ORION.svg?token=jz2ngWM3JpnFiWRYz9Ru&branch=master)](https://travis-ci.com/MaximeNdutiye/ORION)

### Building the Docker images
A Dockerfile is provided to build a container with the required dependencies.

`make container`

### Set up AWS-CLI
Set up AWS credentials in the container using make and a **credentials.csv**
file containing credentials of an IAM user capable of deploying.

Run `make setupAWS` after mounting the **.csv** inside the container in `/go`

### Running the container
Run and mount the needed dirrectories into the container

`make runcontainer`

### Building the lamda
The lambda function must be package into a zip file so it can be deployed by terraform
`make zip`

### Run Tests
`make test`

### FUTURE
Things that we want to get done

### Built with :purple_heart: By:

[Maxime](https://github.com/MaximeNdutiye)