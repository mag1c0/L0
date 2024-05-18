FROM golang:1.22.2-alpine AS builder

COPY . /github.com/mag1c0/L0/backend/source/
WORKDIR /github.com/mag1c0/L0/backend/source/

RUN go mod download
RUN go build -o ./bin/backend cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/mag1c0/L0/backend/source/config ./config
COPY --from=builder /github.com/mag1c0/L0/backend/source/bin/backend .

CMD ["./backend"]