FROM golang:1.22

WORKDIR /app

# goのホットリロードツールをインストール
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# airでアプリケーションを起動
CMD ["air"]

# Cloud Runにデプロイする時はairをダウンロードせず以下を使用。
# TODO: github actionで自動化

# RUN go build -o main .

# EXPOSE 8080

# CMD ["./main"]