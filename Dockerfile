FROM golang:1.20.4 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags dynamic -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest

COPY --from=build /app/main .
COPY --from=build /app/.env .

EXPOSE 9080
CMD ["./main"]
