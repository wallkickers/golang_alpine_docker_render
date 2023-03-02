# ベースとなるDockerイメージ指定
FROM golang:1.20.1-alpine3.17 as dev

# 環境変数
ENV ROOT=/go/src/app
ENV CGO_ENABLED=0

# コンテナログイン時のディレクトリ指定
WORKDIR ${ROOT}

# 必要なライブラリのインストール
RUN apk update && apk --no-cache add bash go git ca-certificates curl vim

# go.mod
# COPY go.mod go.sum ./
# RUN go mod download

# TZ,言語は未設定

COPY ./app ${ROOT}

EXPOSE 3010

CMD ["go", "run", "main.go"]

