FROM golang:1.20 AS builder

WORKDIR /app

COPY source dest

RUN go mod tidy

COPY source dest

RUN go build -o scraper

FROM debian:bullseye-slim

RUN apt-get uppdate && apt-get install -y ca-certificates && apt-get clean

COPY --from=builder /app/scraper /usr/local/bin/scraper

CMD ["scraper"]
