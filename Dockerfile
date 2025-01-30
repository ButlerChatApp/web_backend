FROM golang:1.23.4

WORKDIR /app

# goのホットリロードツールをインストール
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# airでアプリケーションを起動
CMD ["air"]