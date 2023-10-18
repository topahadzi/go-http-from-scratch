
FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]