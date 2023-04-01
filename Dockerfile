FROM golang:1.18-alpine as base

RUN apk add build-base

WORKDIR "/app"

ENV go env -w GO111MODULE=on
ENV CGP_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build ./... && go build

#Second stage building

FROM alpine
RUN apk update upgrade
WORKDIR /app

COPY --from=base /app .
RUN chmod +x go-authentication-api

EXPOSE 8081

CMD ["./go-authentication-api"]

