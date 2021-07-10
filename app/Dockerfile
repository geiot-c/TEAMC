FROM golang:latest

# WORKDIR /go/src/Junking/views 
# RUN apt-get update && apt-get install -y git && apt-get install -y nodejs npm && npm install
WORKDIR /go/src/TEAMC

ENV GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "main.go"]