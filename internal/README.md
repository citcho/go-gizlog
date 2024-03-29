# ディレクトリ構成

```
└── app
    ├── domain (ドメイン層)
    ├── infrastructure (インフラストラクチャ層)
    ├── interfaces (ユーザインターフェース/プレゼンテーション層)
    └── usecases (ユースケース層)
```

## ドメイン層（Enterprise Business Rules）
- entity
真のビジネス領域を型で定義し、NewXXX（ファクトリ）を提供

- value object
エンティティの各属性のルールを定義

- domain service
ファクトリ、リポジトリを呼び出す場所
メソッド名はユビキタス言語を表すようにし、状態を持たない。

- repository


## ユースケース層（Application Business Rules）
- このアプリで必要となる手続きを順番に記述するべきレイヤーで、基本的にドメイン層にメソッド等の形で定義されたビジネスルールを順番に呼び出すことだけが仕事。
- もしこの層で計算や条件判定を書き始めたら、それはドメインオブジェクト自体に定義すべきビジネスロジックではないか？と検討する。
- もし業務手続きが複数のドメインの集約に跨って一つのドメインエンティティ内に実装してしまうとかえって単一責任の原則に違反することになる場合は、ドメインサービスにロジックを定義して、この層から呼び出すことを検討する。
- システムが無くなっても業務で必要になるロジックはドメイン層、業務をシステム化したから必要になったロジックはアプリケーションサービス層に実装すべき。

## ユーザインターフェース/プレゼンテーション層（Interface Adapters）
>ユーザに情報を表示して、ユーザのコマンドを解釈する責務を負う。外部アクタは人間のユーザではなく、別のコンピュータシステムのこともある。

- Controller
データを受け取りDTOに変換してUseCaseに渡す

- Presenter

- Gateway

- middleware(CORS/authentication)

## インフラストラクチャ層（Frameworks & Drivers）

・handler
・third package
