# build stage
FROM golang:1.20.6-alpine3.18 AS builder

WORKDIR /app
ADD . .
RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
RUN GOPROXY=https://goproxy.cn,direct go build -o main main.go && chmod +x main

# run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY app.env .
COPY sqlc.yaml .
COPY db/migration ./db/migration
COPY start.sh .
COPY wait-for .

EXPOSE 8090

ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
