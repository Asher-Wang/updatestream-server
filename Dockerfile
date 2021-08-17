# syntax=docker/dockerfile:experimental

# build state
FROM golang:alpine AS builder
RUN apk update && apk upgrade && apk add --no-cache make gcc git libc-dev openssh-client ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir ~/.ssh /app
RUN ssh-keyscan github.com >> ~/.ssh/known_hosts
RUN git config --global url."git@github.com:hotstar".insteadOf "https://github.com/hotstar"

ENV GOPATH=/go
COPY dummy-domain-service /go/src/github.com/hotstar/dummy-domain-service
WORKDIR /go/src/github.com/hotstar/dummy-domain-service
RUN --mount=type=ssh,id=github go get -v .
RUN GOOS=linux go build -ldflags "-X main.version=$(git describe --tags --abbrev=0).$(git rev-parse --short HEAD)" -o /app/dummy-domain-service .

FROM alpine
WORKDIR /app
COPY --from=builder /app/dummy-domain-service /app/
ENTRYPOINT ["./dummy-domain-service"]
