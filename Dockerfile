FROM golang:1.24.3 AS builder   

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-api .

# Final stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/product-api /app/product-api
RUN chmod +x /app/product-api
CMD ["/app/product-api"]