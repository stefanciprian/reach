ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /reach
WORKDIR /reach

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o reach

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /reach
WORKDIR /reach
COPY --from=builder /reach/reach .
COPY --from=builder /reach/templates/ templates/
COPY --from=builder /reach/static/ static/

ENV GIN_MODE=release

ENV PORT=1989

EXPOSE $PORT

ENTRYPOINT ["./reach"]