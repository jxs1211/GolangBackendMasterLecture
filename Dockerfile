# build stage
FROM golang:1.20.6-alpine AS builder

WORKDIR /app
ADD . .

RUN GOPROXY=https://goproxy.cn,direct go build -o main main.go && chmod +x main

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY sqlc.yaml .
COPY db/migration ./db/migration

EXPOSE 8090

ENTRYPOINT ["/app/main"]
