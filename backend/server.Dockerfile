FROM golang:latest
RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .


RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD . .
# FROM alpine:latest // multi-stage build
EXPOSE ${SERVER_PORT}

ENTRYPOINT CompileDaemon --build="go build -o simple-user-mgmt main.go" --command=./simple-user-mgmt
