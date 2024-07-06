FROM golang:1.22.5-alpine3.19
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8000
RUN go build -o ./cmd/main/main ./cmd/main
CMD ["./cmd/main/main"]
