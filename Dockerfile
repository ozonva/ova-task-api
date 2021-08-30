# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
RUN apk add --update make

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY Makefile ./

RUN make build

EXPOSE 8080

CMD [ "./bin/main" ]