FROM golang:1.21-alpine AS builder

WORKDIR /build

# Install git for go-git dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=${VERSION:-dev} -X main.commit=${COMMIT:-unknown} -X main.date=${DATE:-unknown}" -o doplan ./cmd/doplan

FROM alpine:latest

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY --from=builder /build/doplan .

ENTRYPOINT ["./doplan"]

