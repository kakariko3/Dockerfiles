# Rails6_API x PostgreSQL Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml
- Gemfile
- Gemfile.lock
- <span>entrypoint.sh</span>

## 2. rails new でアプリを作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm api rails new . --no-deps --database=mysql --api
```
`docker-compose run`コマンドではイメージの構築から、コンテナの構築・起動まで行ってくれる。引数にサービスを指定する必要がある。<br>
このコマンドを実行することで、Dockerfileを元にapiイメージがビルドされ、Railsの各種ファイルが構成される。<br>

`--no-deps` : リンクしたサービスを起動しない<br>
`--database=mysql` : DBにMySQLを指定<br>
`--api` : APIモードでアプリを作成(APIに必要ない部分をデフォルトで作成しなくなる)

## 3. Dockerイメージのビルド

先ほどの`rails new`により、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドを実行することで、Dockerイメージををビルドする際に`bundle install`が行われる。
```
docker-compose build
```

## 4. database.yml の設定と、DBの作成

`rails new`で生成された`config/database.yml`を下記のように書き換える。
```
default: &default
  adapter: mysql2
  encoding: utf8mb4
  pool: <%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %>
  username: <%= ENV.fetch("MYSQL_USERNAME", "root") %>
  password: <%= ENV.fetch("MYSQL_PASSWORD", "password") %>
  host: <%= ENV.fetch("MYSQL_HOST", "db") %>

development:
  <<: *default
  database: app_development

test:
  <<: *default
  database: app_test

production:
  <<: *default
  database: app_production
  username: app
  password: <%= ENV['APP_DATABASE_PASSWORD'] %>
```

## 5. Dockerコンテナの起動、DBの作成

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up
```
新規ターミナルを開いて下記コマンドを実行し、データベースを作成する。
```
docker-compose run api rails db:create
```
Webブラウザで http://localhost:3000 へアクセスし、Railsが起動していることを確認する。

## 6. その他

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
docker-compose exec <サービス名> bash

# コンテナ内のコマンドを実行
docker-compose run <サービス名> <コマンド>
```

### 参考資料
---
https://blog.cloud-acct.com/posts/u-docker-compose-rails6new/<br>
https://nakatanorihito.com/programming/docker-rails-postgresql
