FROM golang:1.19.5-alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o main .

EXPOSE $PORT

CMD ["./main"]
