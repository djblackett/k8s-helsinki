FROM golang:latest as builder
LABEL authors="dave"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
    go install github.com/jessevdk/go-assets-builder@latest
COPY *.go ./

RUN go-assets-builder templates -o assets.go && \
    CGO_ENABLED=0 GOOS=linux go build -o /todo-go

FROM scratch

COPY --from=builder /todo-go .
CMD ["/todo-go"]