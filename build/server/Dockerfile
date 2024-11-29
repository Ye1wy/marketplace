FROM golang:1.20 AS builder

WORKDIR /app

COPY source dest

RUN go mod tidy

COPY source dest

RUN go build -o server .

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

COPY --from=builder /app/server usr/local/bin/server

CMD [ "server" ]
