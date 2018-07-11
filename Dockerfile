# Use apline base image
FROM alpine:3.6

# Configure Environment
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV HOME /home

# AWS-CLI
RUN apk --update --no-cache add groff less bash python py-pip go musl-dev make zip && \
    pip install 'awscli>=1.11.109' && \
    apk --purge -v del py-pip

# Terraform, git, curl, go
RUN apk add --update --no-cache git curl && \
    curl https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip > terraform_0.11.7_linux_amd64.zip && \
    unzip terraform_0.11.7_linux_amd64.zip -d /bin && \
    rm -f terraform_0.11.7_linux_amd64.zip

RUN mkdir $GOPATH
ADD makefile $GOPATH/makefile
WORKDIR $GOPATH