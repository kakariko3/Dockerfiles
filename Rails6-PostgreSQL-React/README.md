# Rails6 x PostgreSQL x React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、その直下に下記のようにファイルを配置する。
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

## 2. rails new でアプリ作成

作業ディレクトリに移動し、下記コマンドを実行する。
```
$ docker-compose run api rails new . --force --no-deps --database=postgresql --api
```

## 3. Dockerイメージのビルド

`rails new` で各種ファイルを作成したら、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドで、Dockerイメージををビルドする際に `bundle install` を行ってくれる。
```
$ docker-compose build
```

## 4. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
$ docker-compose run --rm front sh -c "npm install -g create-react-app && create-react-app reactapp"
```

## 5. database.yml の設定

`rails new` で作成された `api/config/database.yml` を下記のように書き換える。
```
default: &default
  adapter: postgresql
  encoding: unicode
  host: db
  username: postgres
  password:
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
下記コマンドを実行し、データベースを作成する。
```
$ docker-compose run api rails db:create
```
Webブラウザを起動して以下にアクセスし、http://localhost:3000 でRails、http://localhost:8000 でReactが起動していることを確認する。

### 参考資料
---
https://qiita.com/dl10yr/items/b76969da1c2c33595a4a<br>
https://usconsort.com/docker-env-vol1/
