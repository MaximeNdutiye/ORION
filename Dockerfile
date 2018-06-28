# Use an official Python runtime as a parent image
FROM alpine:3.6

# Enviroment variables
ENV GOPATH /home
ENV HOME /home

# Variables
ENV TERRAFORM_VERSION = 0.11.7
ENV GO_VERSION = 10.3

# AWS-CLI installation
RUN apk --update --no-cache add groff less bash python py-pip && \
    pip install 'awscli>=1.11.109' && \
    pip install 'awscli-cwlogs' && \
    apk --purge -v del py-pip

# Install Terraform, git, curl
RUN apk add --update --no-cache git curl && \
    curl https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip > terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /bin && \
    rm -f terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Install go
RUN curl -O https://dl.google.com/go/go1.${GO_VERSION}.linux-amd64.tar.gz && \
    tar -xvf go1.${GO_VERSION}.linux-amd64.tar.gz && mv go /usr/local && \
    export PATH=$PATH:/usr/local/go/bin

# Set working directory to home
WORKDIR $HOME

# Copy the current directory contents into the container at /app
ADD . $HOME

# Make port 80 available to the world outside this container
EXPOSE 80
