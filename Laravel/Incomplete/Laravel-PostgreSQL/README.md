# Laravel x PostgreSQL

## 1. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build --no-cache
```

## 2. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```

## 3. Laravelプロジェクトの作成

下記コマンドを実行し、appコンテナ上でLaravelプロジェクトのベースを作成する。
```
docker-compose exec app composer create-project --prefer-dist laravel/laravel my-app
```
Webブラウザで http://localhost:8000 へアクセスし、Laravelが起動していることを確認する。

## 参考資料

### Laravel x MySQL

https://maasaablog.com/development/laravel/758/<br>
https://qiita.com/ucan-lab/items/5fc1281cd8076c8ac9f4

### Laravel x PostgreSQL

https://qiita.com/sakeafterbeer/items/56cea7e981dacdfc686f<br>
https://zenn.dev/nagi125/articles/ea1d314c94409341a3b0<br>
https://lightning-shine.com/post1437/

### Laravel x React

https://maasaablog.com/development/laravel/891/
