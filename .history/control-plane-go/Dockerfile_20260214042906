# -------- Build Stage --------
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o control-plane cmd/server/main.go


# -------- Runtime Stage --------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/control-plane .

EXPOSE 8080

CMD ["./control-plane"]
