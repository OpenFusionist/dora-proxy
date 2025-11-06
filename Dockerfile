# syntax=docker/dockerfile:1

ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /out/dora-proxy .

FROM alpine:${ALPINE_VERSION}

RUN apk add --no-cache ca-certificates tzdata && adduser -D -H -u 10001 appuser

WORKDIR /app

COPY --from=builder /out/dora-proxy /app/dora-proxy

EXPOSE 8081

USER appuser

ENV PROXY_LISTEN_ADDR=:8081 \
    PROXY_UPSTREAM_BASE_URL=http://localhost:8080 \
    PROXY_CONSENSUS_API_URL=http://localhost:5052

ENTRYPOINT ["/app/dora-proxy"]


