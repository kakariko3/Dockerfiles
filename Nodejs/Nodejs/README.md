# Setup Node.js

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── app
│   ├── .gitignore
│   └── Dockerfile
├── .gitignore
└── docker-compose.yml
```

## 2. Dockerコンテナの起動

作業ディレクトリに移動し、下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```

下記コマンドを実行し、appコンテナに入る。
```
docker-compose exec app ash
```

## 3. その他

## 参考資料
