FROM golang:latest as builder
LABEL authors="dave"

WORKDIR /app


COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /add-reminder

FROM alpine:latest

COPY --from=builder /add-reminder .

CMD ["/add-reminder"]