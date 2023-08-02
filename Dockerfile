FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM scratch

COPY --from=builder /app/myapp /myapp

ENV PORT=8080

EXPOSE $PORT

CMD ["/myapp"]
