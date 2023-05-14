FROM golang:1.20.4 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

FROM alpine:latest

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 9080
CMD ["./main"]
