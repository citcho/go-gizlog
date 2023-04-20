# バイナリ作成用コンテナステージ
FROM golang:1.20.3-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app

# ------------------------------------------------------------

# デプロイ用コンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ------------------------------------------------------------

# ローカル用ライブリロード対応コンテナステージ
FROM golang:1.20.3 as dev

WORKDIR /app

RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/rogpeppe/godef@latest \
    && go install github.com/cosmtrek/air@latest


WORKDIR /app/cmd/

CMD ["air"]
