FROM golang:1.19.5-alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]
