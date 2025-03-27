ディレクトリ構成(クリーンアーキテクチャ)
```myapp/
    ├── cmd/             # 実行可能ファイルのエントリーポイント
    │   ├── server/      # サーバー用のエントリーポイント
    │   │   ├── main.go
    │   ├── worker/      # バックグラウンドジョブ用
    │   │   ├── main.go
    │   ├── migrate/     # データベースマイグレーション用
    │   │   ├── main.go
    ├── config/          # 設定関連
    │   ├── config.go
    ├── internal/        # アプリケーション固有のロジック
    │   ├── domain/      # エンティティ（ビジネスルール）
    │   │   ├── user.go
    │   ├── repository/  # データベース操作
    │   │   ├── user_repository.go
    │   ├── usecase/     # ユースケース（ビジネスロジック）
    │   │   ├── user_usecase.go
    │   ├── service/     # サービス（補助的な処理）
    │   │   ├── email_service.go
    │   ├── handler/     # API ハンドラ
    │   │   ├── user_handler.go
    │   ├── middleware/  # ミドルウェア
    │   │   ├── auth_middleware.go
    │   ├── validator/   # 入力バリデーション
    │   │   ├── user_validator.go
    │   ├── router/      # ルーティング管理
    │   │   ├── router.go
    ├── pkg/             # 外部に公開可能なライブラリ
    │   ├── logger/      # ログ関連
    │   │   ├── logger.go
    │   ├── database/    # データベース関連
    │   │   ├── db.go
    │   ├── cache/       # キャッシュ（Redis）
    │   │   ├── redis.go
    │   ├── queue/       # メッセージキュー（Kafka, RabbitMQ）
    │   │   ├── queue.go
    │   ├── httpclient/  # 外部 API 呼び出し
    │   │   ├── client.go
    ├── api/             # API 定義（OpenAPI や Protobuf）
    │   ├── openapi.yaml
    ├── migrations/      # データベースマイグレーションファイル
    │   ├── 20240325_create_users_table.up.sql
    │   ├── 20240325_create_users_table.down.sql
    ├── test/            # 統合テスト
    │   ├── integration_test.go
    ├── docs/            # API ドキュメント管理
    │   ├── README.md
    │   ├── architecture.md
    ├── scripts/         # シェルスクリプト管理
    │   ├── build.sh
    │   ├── deploy.sh
    ├── assets/          # 画像や静的ファイル
    │   ├── images/
    ├── web/             # フロントエンド (Next.js や React)
    │   ├── package.json
    │   ├── src/
    ├── .github/
```