FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM debian:buster-slim

COPY --from=builder /app/myapp /usr/local/bin/myapp

CMD ["myapp"]
