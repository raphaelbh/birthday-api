FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM golang:1.20

COPY --from=builder /app/myapp /usr/local/bin/myapp

EXPOSE 80

CMD ["myapp"]
