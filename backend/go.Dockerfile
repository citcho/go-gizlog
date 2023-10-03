# バイナリ作成用コンテナステージ
FROM golang:1.21.0-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go build -trimpath -ldflags "-w -s" -o app

# WORKDIR /app/cmd/monolith
# RUN go build -trimpath -ldflags "-w -s" -o ../../app
# WORKDIR /app/cmd/migrate
# RUN go build -trimpath -ldflags "-w -s" -o ../../migrate

# ------------------------------------------------------------

# デプロイ用コンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ------------------------------------------------------------

# ローカル用ライブリロード対応コンテナステージ
FROM golang:1.21.0 as dev

WORKDIR /app

RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/rogpeppe/godef@latest \
    && go install github.com/golang/mock/mockgen@v1.6.0 \
    && go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "./.air.toml"]
