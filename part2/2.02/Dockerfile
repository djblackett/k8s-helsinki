FROM golang:latest as builder
LABEL authors="dave"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY templates ./templates
COPY todo-react/build ./build

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-go

FROM alpine:latest

COPY --from=builder /todo-go .
COPY --from=builder /app/templates /templates
COPY --from=builder /app/build /build
RUN mkdir tmp/kube

CMD ["/todo-go"]