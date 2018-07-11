# ORION
Serverless image scaling service built to run in AWS.

[![Build Status](https://travis-ci.com/MaximeNdutiye/ORION.svg?token=jz2ngWM3JpnFiWRYz9Ru&branch=master)](https://travis-ci.com/MaximeNdutiye/ORION)

### Building the Docker images
Building the docker container will create and configure a container with everything needed for development.
The docker container is not needed, but it is useful since the environement is preconfigured.
`Note`: If the container isn't being used then the correct version of terraform and other tools must be installed by the user

`container`

### Set up AWS-CLI
AWS credentials must be setup using aws-cli inside the container to be able to deploy to aws
After creating a user with credentials to deploy download the .csv containing the user credentials from aws and then run
`make setupAWS`

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