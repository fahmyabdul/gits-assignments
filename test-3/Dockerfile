FROM golang:1.19.2-bullseye AS builder

ENV GO111MODULE="auto"

RUN apt-get install git

WORKDIR /app/test-3-grpc

COPY . .

RUN go mod vendor

FROM golang:1.19.2-bullseye

ENV GO111MODULE="auto"

WORKDIR /app/test-3-grpc

COPY --from=builder /app/test-3-grpc ./

# RUN mv ./.config.dev.yml ./.config.yml

EXPOSE 11000

CMD ["make", "run-grpc"]