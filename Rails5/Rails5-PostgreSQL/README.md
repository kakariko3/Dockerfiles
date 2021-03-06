# Rails5 x PostgreSQL Setup

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
$ docker-compose run web rails new . --force --no-deps --database=postgresql --skip-bundle
```
`docker-compose run`コマンドではイメージの構築から、コンテナの構築・起動まで行ってくれる。引数にサービスを指定する必要がある。<br>
このコマンドを実行することで、Dockerfileを元にwebイメージがビルドされ、Railsの各種ファイルが構成される。<br>

`--force` : 既存のGemfileを上書きするためのオプション<br>
`--no-deps` : リンクしたサービスを起動しない<br>
`--database=postgresql` : DBにPostgreSQLを指定<br>
`--skip-bundle` : bundleをスキップ(最後にbundleする)

## 3. Dockerイメージのビルド

先ほどの`rails new`により、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドを実行することで、Dockerイメージををビルドする際に`bundle install`が行われる。
```
$ docker-compose build
```

## 4. database.yml の設定と、DBの作成

`rails new`で生成された`config/database.yml`を下記のように書き換える。
```
default: &default
  adapter: postgresql
  encoding: unicode
  host: db
  username: postgres
  password: password
  pool: 5

development:
  <<: *default
  database: myapp_development

test:
  <<: *default
  database: myapp_test
```

## 5. Dockerコンテナの起動、DBの作成

下記コマンドを実行し、コンテナを起動する。
```
$ docker-compose up
```
新規ターミナルを開いて下記コマンドを実行し、データベースを作成する。
```
$ docker-compose run web rails db:create
```
Webブラウザで http://localhost:3000 へアクセスし、Railsが起動していることを確認する。

## 6. その他

### dockerコマンド
```
# コンテナ一覧の表示
$ docker ps -a

# イメージ一覧
$ docker images -a

# 停止中のコンテナを削除
$ docker container prune

# <none>タグのイメージを一括削除
$ docker image prune
```

### docker-composeコマンド
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
https://qiita.com/daichi41/items/dfea6195cbb7b24f3419<br>
https://docs.docker.com/samples/rails/
