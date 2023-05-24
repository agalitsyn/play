FROM golang:latest

WORKDIR /app

COPY go.* /app/
COPY *.go /app/

RUN go build -o app file.go

CMD [ "/app/app" ]
