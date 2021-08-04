FROM golang:1.16

WORKDIR /app

 
# RUN --mount=type=secret,id=github_token echo "[url \"https://$(cat /run/secrets/github_token)@github.com/hotstar\"]\n\tinsteadOf = https://github.com/hotstar" >> /root/.gitconfig
RUN --mount=type=secret,id=github_token git config --global url."https://$(cat /run/secrets/github_token)@github.com/hotstar".insteadOf "https://github.com/hotstar"
RUN cat /root/.gitconfig

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build server.go

ENTRYPOINT [ "/app/server" ]