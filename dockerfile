FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd

EXPOSE 8080

CMD ["./auth-service"]
