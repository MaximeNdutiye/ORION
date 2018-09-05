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

or manually setup credentials yourself using `aws configure` inside the container

or simply setup environment variables

```
$ export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
$ export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
$ export AWS_DEFAULT_REGION=us-west-2
```

### Running the container
Run the container with the required directories using the following command

`make runcontainer`

### Building & Packaging
The lambda function must be package into a zip file so it can be deployed by terraform
`make zip` will build an **orion** executable and a **lambda.zip**

### Deploying & Destroying
`make deploy` will automatically build **lambda.zip** and use terraform to create the AWS architecure

`make destroy` conversly destroys AWS deployment 

### Run Tests
Run a simple test to check the API Gateway endpoint to ensure that the lambda is working as intended
`make test`

### AWS Architecture Diagram
![architecture.png](https://github.com/MaximeNdutiye/ORION/blob/master/images/architecture.png)

### FUTURE
- Require authentication to access lambda
- More features: filters, convertion, video scaling...
- Seamless upgrades and rollbacks by versioning resources with prefixes

### Built with :purple_heart: By:

[Maxime](https://github.com/MaximeNdutiye)