# Rails6 x MySQL Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、その直下に下記のファイルを移動する。
- Dockerfile
- docker-compose.yml
- Gemfile
- Gemfile.lock
- <span>entrypoint.sh</span>

## 2. rails new でアプリ作成

作業ディレクトリにいることを確認し、下記コマンドを実行する。
```
$ docker-compose run web rails new . --force --no-deps --database=mysql --skip-test --webpacker
```
後にRSpecを導入するため、`--skip-test` オプションでMinitestが作成されないようにする。<br>
`--webpacker` オプションでwebpackerをインストールする。

## 3. Dockerイメージのビルド

`rails new` で各種ファイルの作成、webpackerのインストールが完了したら、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドを実行すると、Dockerイメージををビルドする際に `bundle install` を行ってくれる。
```
$ docker-compose build
```

## 4. database.yml の設定

`rails new` で作成された `config/database.yml` を下記のように書き換える。
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
  database: myapp_development

test:
  <<: *default
  database: myapp_test

production:
  <<: *default
  database: myapp_production
  username: myapp
  password: <%= ENV['MYAPP_DATABASE_PASSWORD'] %>
```

## 5. Dockerコンテナの起動、DBの作成

下記コマンドを実行し、コンテナを起動する。
```
$ docker-compose up
```
新規のターミナルで下記コマンドを実行し、データベースを作成する。
```
$ docker-compose run web rails db:create
```
Webブラウザで http://localhost:3000 へアクセスし、Railsが起動していることを確認する。

## 6. その他

### dockerコマンド
---
```
# コンテナ一覧の表示
$ docker ps -a

# イメージ一覧
$ docker images -a

# <none>タグのイメージを一括削除
$ docker image prune -a
```

### docker-composeコマンド
---
```
# 起動
$ docker-compose up

# バックグラウンドで起動
$ docker-compose up -d

# 停止
$ docker-compose stop

# 停止＆削除
$ docker-compose down

# 稼働中のコンテナに入る
$ docker-compose exec <サービス名> bash

# コンテナ内のコマンドを実行
$ docker-compose run <サービス名> <コマンド>
```

### 参考資料
---
https://qiita.com/nsy_13/items/9fbc929f173984c30b5d<br>
https://qiita.com/kodai_0122/items/795438d738386c2c1966
