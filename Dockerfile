FROM golang:1.25.1-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apk update && apk add make
RUN make build

FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /app/bin/moviemanager_echo/main .
COPY --from=builder /app/certs ./certs

EXPOSE 7734

ENV GMM_CERT_PATH=certs/myCA.crt
ENV GMM_KEY_PATH=certs/myCA.key
CMD ["./main"]
