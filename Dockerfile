FROM ubuntu:18.04

ENV GO_VERSION=1.9
ENV TF_VERSION=terraform_0.11.7
ENV GOROOT /goroot
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# curl make zip
RUN apt-get update && apt-get install -y curl make zip

# AWS-CLI
RUN apt-get install -y python-pip python && \
    pip install 'awscli>=1.11.109'

# Go
RUN mkdir $GOROOT && \
    curl \
    https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz \
    | tar xvzf - -C $GOROOT --strip-components=1 && \
    mkdir $GOPATH && \
    apt-get clean all

# Terraform
RUN curl \
    https://releases.hashicorp.com/terraform/0.11.7/${TF_VERSION}_linux_amd64.zip \
    > ${TF_VERSION}_linux_amd64.zip && \
    unzip ${TF_VERSION}_linux_amd64.zip -d /bin && \
    rm -f ${TF_VERSION}_linux_amd64.zip

ADD makefile $GOPATH/makefile
WORKDIR $GOPATH