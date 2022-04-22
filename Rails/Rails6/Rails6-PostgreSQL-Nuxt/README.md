# Setup Rails6 x PostgreSQL x Nuxt.js

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend
│   ├── Dockerfile
│   ├── entrypoint.sh
│   ├── Gemfile
│   └── Gemfile.lock
├── frontend
│   └── Dockerfile
├── .gitignore
└── docker-compose.yml
```

## 2. Railsアプリを作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm backend rails new . --force --skip-bundle --database=postgresql --api
```
`docker-compose run`コマンドではイメージの構築から、コンテナの構築・起動までを行う。引数にサービスを指定する必要がある。<br>
このコマンドを実行することで、Dockerfileを元にbackendイメージがビルドされ、Railsの各種ファイルが構成される。<br>

`--force` : 既存のGemfileを上書きする<br>
`--skip-bundle` : bundle installを実行しない<br>
`--database=postgresql` : DBにPostgreSQLを指定<br>
`--api` : APIモードでアプリを作成(APIに必要ない部分をデフォルトで作成しなくなる)

## 3. Dockerイメージのビルド

下記コマンドを実行し、Dockerイメージのビルドを行う。
```
docker-compose build --no-cache
```
先ほどの`rails new`によりGemfileが更新されているため、backendイメージは再ビルドされ、`bundle install`が行われる。

## 4. Nuxt.jsアプリを作成

下記コマンドを実行し、Nuxt.jsアプリを作成する。
```
docker-compose run --rm frontend npx create-nuxt-app app
```
下記URLを参考に、create-nuxt-appの各オプションを選択する。<br>
https://blog.proglus.jp/4972/

下記コマンドを実行し、Reactアプリをappディレクトリからfrontendディレクトリ直下に移動する。
```
mv frontend/app/{*,.*} frontend
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル

下記コマンドを実行し、空になったappディレクトリを削除する。
```
rmdir frontend/app
```

## 5. database.yml の設定

`rails new`で生成された`backend/config/database.yml`を下記のように書き換える。
```
default: &default
  adapter: postgresql
  encoding: unicode
  pool: <%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %>
  username: <%= ENV.fetch("POSTGRES_USERNAME", "postgres") %>
  password: <%= ENV.fetch("POSTGRES_PASSWORD", "postgres") %>
  host: <%= ENV.fetch("POSTGRES_HOST", "db") %>

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

また、`backend/.gitignore`に下記のテキストを追記する。
```
# Original
.DS_Store
/.ash_history
/.irb_history
```

## 6. DBの作成、Dockerコンテナの起動

下記コマンドを実行し、データベースを作成する。
```
docker-compose run --rm backend rails db:create
```
下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザを起動して以下にアクセスし、http://localhost:3000 でRails、http://localhost:8080 でNuxt.jsが起動していることを確認する。

## 7. その他

## 参考資料

https://blog.cloud-acct.com/posts/u-docker-compose-rails6new/<br>
https://blog.cloud-acct.com/posts/u-docker-create-nuxtjs/<br>
https://zenn.dev/kkosuke/articles/c3f129195a07ad<br>
https://qiita.com/at-946/items/c69a512ea47941747b18
