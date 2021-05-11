# Rails6 x PostgreSQL x React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
api
  - Dockerfile
  - entrypoint.sh
  - Gemfile
  - Gemfile.lock
front
  - Dockerfile
docker-compose.yml
```

## 2. rails new でアプリを作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
$ docker-compose run api rails new . --force --no-deps --database=postgresql --api
```
`docker-compose run`コマンドではイメージの構築から、コンテナの構築・起動まで行ってくれる。引数にサービスを指定する必要がある。<br>
このコマンドを実行することで、Dockerfileを元にapiイメージがビルドされ、Railsの各種ファイルが構成される。<br>

`--force` : 既存のGemfileを上書きするためのオプション<br>
`--no-deps` : リンクしたサービスを起動しない<br>
`--database=postgresql` : DBにPostgreSQLを指定<br>
`--api` : APIモードでアプリを作成(APIに必要ない部分をデフォルトで作成しなくなる)

## 3. Dockerイメージのビルド

先ほどの`rails new`により、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドを実行することで、Dockerイメージををビルドする際に`bundle install`が行われる。
```
$ docker-compose build
```

## 4. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
$ docker-compose run --rm front npx create-react-app react_app --template typescript
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。

## 5. database.yml の設定

`rails new`で生成された`api/config/database.yml`を下記のように書き換える。
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

## 6. Dockerコンテナの起動、DBの作成

下記コマンドを実行し、コンテナを起動する。
```
$ docker-compose up
```
新規ターミナルを開いて下記コマンドを実行し、データベースを作成する。
```
$ docker-compose run api rails db:create
```
Webブラウザを起動して以下にアクセスし、http://localhost:3000 でRails、http://localhost:8000 でReactが起動していることを確認する。

## 7. その他

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
https://qiita.com/dl10yr/items/b76969da1c2c33595a4a<br>
https://usconsort.com/docker-env-vol1/
