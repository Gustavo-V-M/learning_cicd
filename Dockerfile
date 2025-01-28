FROM golang:1.23

EXPOSE 8080/tcp

WORKDIR /usr/src/app

COPY ./src/go.mod ./
RUN go mod download && go mod verify

COPY ./src/. ./
RUN go build -v http_example .
ENTRYPOINT ["./http_example"]
