FROM golang:latest AS build

WORKDIR /go/src/
ENV GO111MODULE=on
COPY go.* .
RUN go mod download && go mod tidy
COPY . .

RUN go build -o /go/bin/app && \
	go install github.com/cosmtrek/air@latest

CMD ["air"]
