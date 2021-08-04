FROM golang:1.16

WORKDIR /app

RUN git config --global url."git@github.com:hotstar".insteadOf "https://github.com/hotstar"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build server.go

ENTRYPOINT [ "/app/server" ]