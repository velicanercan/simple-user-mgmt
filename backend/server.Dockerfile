FROM golang:latest
RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .


RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD . .
EXPOSE ${SERVER_PORT}

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
