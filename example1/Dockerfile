# Using ubuntu image like a base

FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y curl apt-utils wget net-tools unzip

RUN mkdir -p /downloads
WORKDIR /downloads

#install golang
RUN wget -nv https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
RUN ls -la /downloads/go1.18.3.linux-amd64.tar.gz
RUN mkdir -p /usr/local
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN go version

#gRPC
RUN wget -nv https://github.com/protocolbuffers/protobuf/releases/download/v21.1/protoc-21.1-linux-x86_64.zip
RUN mkdir -p /protoc
WORKDIR /protoc
RUN unzip /downloads/protoc-21.1-linux-x86_64.zip
ENV PATH="$PATH:/protoc/bin/"
RUN protoc --version
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ENV PATH="$PATH:/root/go/bin/"

#ADD content/go.mod .
#ENV GO111MODULE=on
#RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
#ENV PATH="$PATH:$(go env GOPATH)/bin"


