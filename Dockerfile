
FROM golang:1.12.5-alpine3.9 as builder
COPY main.go .
COPY static/* /app/static/
RUN go build -o /app/main main.go

FROM alpine:3.9
WORKDIR /app
COPY --from=builder /app .
CMD ["./main"]
