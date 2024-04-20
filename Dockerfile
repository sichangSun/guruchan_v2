
FROM golang:1.17.0-alpine3.13 AS builder

WORKDIR /app
COPY ./main.go ./

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o main main.go

FROM alpine:3.13
WORKDIR /app
# Build stage からビルドされた main だけを Run stage にコピーする。
COPY --from=builder /app/main .
# local env copy
# COPY .env .

CMD ["./main"]
