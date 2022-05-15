FROM golang:1.18-alpine

RUN apk add build-base

WORKDIR "/app"

ENV go env -w GO111MODULE=on
ENV CGP_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .
COPY . .

RUN go build ./... && go build

EXPOSE 8081

CMD ["./go-authentication-api"]

