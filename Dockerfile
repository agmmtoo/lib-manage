FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN apk update

RUN go build -o ./main ./cmd/libraryapp

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"]