# Next.js Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build
```

## 3. Reactのインストール + アプリ作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm front npx create-next-app next_app
```
`--rm` : 停止後コンテナを削除<br>
`--use-npm` : パッケージマネージャーにnpmを指定<br>
`--example xxx` : テンプレートを指定

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up
```
Webブラウザで http://localhost:8000 へアクセスし、Next.jsが起動していることを確認する。

## 5. その他

### dockerコマンド
```
# コンテナ一覧の表示
docker ps -a

# イメージ一覧
docker images -a

# 停止中のコンテナを削除
docker container prune

# <none>タグのイメージを一括削除
docker image prune

# コンテナ、ボリューム、ネットワーク、イメージを一括削除
docker system prune -a
```

### docker-composeコマンド
```
# コンテナの作成と起動
docker-compose up

# バックグラウンドで起動
docker-compose up -d

# 起動
docker-compose start

# 停止
docker-compose stop

# 停止＆削除
docker-compose down

# コンテナのログを表示
docker-compose logs -f

# 稼働中のコンテナに入る
docker-compose exec <サービス名> ash

# コンテナ内のコマンドを実行
docker-compose run <サービス名> <コマンド>
```

### 参考資料
---
https://zenn.dev/taku1115/articles/6c9fa97ab37e38
