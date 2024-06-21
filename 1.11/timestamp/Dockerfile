FROM golang:latest as builder
LABEL authors="djblackett"

WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /timestamp

FROM alpine:latest

COPY --from=builder /timestamp .
CMD ["/timestamp"]