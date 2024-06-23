FROM golang:latest as builder
LABEL authors="dave"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY templates ./templates

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-go

FROM scratch

COPY --from=builder /todo-go .
COPY --from=builder /app/templates /templates
CMD ["/todo-go"]