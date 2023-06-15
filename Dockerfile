FROM golang:1.19

WORKDIR /order-microservice

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./cmd ./...

CMD ["go","run","./cmd"]