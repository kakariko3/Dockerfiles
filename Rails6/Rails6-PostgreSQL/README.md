# Rails6 x PostgreSQL Setup

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
docker-compose run app rails new . --force --no-deps --database=postgresql --skip-bundle
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
docker-compose build
```

## 4. database.yml の設定と、DBの作成

`rails new`で生成された`config/database.yml`を下記のように書き換える。
```
default: &default
  adapter: postgresql
  encoding: unicode
  pool: <%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %>
  username: <%= ENV.fetch("POSTGRES_USERNAME", "postgres") %>
  password: <%= ENV.fetch("POSTGRES_PASSWORD", "password") %>
  host: <%= ENV.fetch("POSTGRES_HOST", "db") %>

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
docker-compose up
```
新規ターミナルを開いて下記コマンドを実行し、データベースを作成する。
```
docker-compose run app rails db:create
```
Webブラウザで http://localhost:3000 へアクセスし、Railsが起動していることを確認する。

## 6. その他

## エラーが起きた場合の対処

下記のようなエラーが表示され、ターミナルが止まってしまった場合、
```
myapp_web_1 exited with code 1
```
コンテナを一度終了させてから、Webpackerをインストールするコマンドを実行する。
```
docker-compose down
```
```
docker-compose run app bundle exec rails webpacker:install
```
再度コンテナを起動する。
```
docker-compose up
```

## 参考資料

https://qiita.com/kodai_0122/items/795438d738386c2c1966<br>
https://qiita.com/me-654393/items/ac6f61f3eee66380ecd7<br>
https://docs.docker.com/samples/rails/
