FROM golang:1.15.6-alpine AS build-env
WORKDIR /app
COPY . .
RUN go build -o goapp

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/goapp .
ENTRYPOINT ./goapp