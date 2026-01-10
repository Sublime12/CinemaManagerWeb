FROM golang:1.25.5-alpine3.23 AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk


RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o cinema_manager_api
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o cinema_manager_seeder ./seeds

FROM alpine:latest

# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/cinema_manager_api .
COPY --from=builder /app/cinema_manager_seeder .

EXPOSE 8080

ENTRYPOINT ["./cinema_manager_api"]
