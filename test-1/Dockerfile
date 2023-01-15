FROM golang:1.19.2-bullseye AS builder

ENV GO111MODULE="auto"

WORKDIR /app

COPY . .

RUN go build -o test-1

FROM golang:1.19.2-bullseye

COPY --from=builder /app/test-1 /usr/bin

ENTRYPOINT ["/usr/bin/test-1"]