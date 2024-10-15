FROM golang:1.23

WORKDIR /go/src/github.com/pirateunclejack/monolith-to-microservice-project
COPY . .

RUN go install github.com/cespare/reflex@latest
