# 環境設定

## protobufのインストール
```sh
brew install protobuf

protoc --version
# libprotoc 3.21.12
```

## Go言語用プロトコルコンパイラのインストール
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"

protoc-gen-go --version
# protoc-gen-go v1.28.1
```
