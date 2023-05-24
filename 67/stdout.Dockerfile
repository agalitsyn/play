FROM golang:latest

WORKDIR /app

COPY go.* /app/
COPY *.go /app/

RUN go build -o app stdout.go

CMD [ "/app/app" ]
