# Setup Rails6 x PostgreSQL x React

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend
│   ├── Dockerfile
│   ├── Gemfile
│   └── Gemfile.lock
├── frontend
│   └── Dockerfile
├── .gitignore
└── docker-compose.yml
```

## 2. rails new でアプリを作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose run --rm backend rails new . --force --skip-bundle --database=postgresql --api
```
`docker-compose run`コマンドではイメージの構築から、コンテナの構築・起動まで行ってくれる。引数にサービスを指定する必要がある。<br>
このコマンドを実行することで、Dockerfileを元にbackendイメージがビルドされ、Railsの各種ファイルが構成される。<br>

`--force` : 既存のGemfileを上書きする<br>
`--skip-bundle` : bundle installを実行しない<br>
`--database=postgresql` : DBにPostgreSQLを指定<br>
`--api` : APIモードでアプリを作成(APIに必要ない部分をデフォルトで作成しなくなる)

## 3. Dockerイメージのビルド

先ほどの`rails new`により、Gemfileが更新されているので、イメージをビルドする。<br>
下記コマンドを実行することで、Dockerイメージををビルドする際に`bundle install`が行われる。
```
docker-compose build --no-cache
```

## 4. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm frontend npx create-react-app app --template typescript
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。
<br>

下記コマンドを実行し、Reactアプリをappディレクトリからfrontendディレクトリ直下に移動する。
```
mv frontend/app/{*,.*} frontend
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル
<br>

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
Webブラウザを起動して以下にアクセスし、http://localhost:8000 でRails、http://localhost:3000 でReactが起動していることを確認する。

## 7. その他

## 参考資料

https://qiita.com/kashimuuuuu/items/b5f35057dfe1980d053a<br>
https://blog.cloud-acct.com/posts/u-docker-compose-rails6new/<br>
https://blog.cloud-acct.com/posts/u-docker-create-nuxtjs/
https://qiita.com/at-946/items/c69a512ea47941747b18<br>
