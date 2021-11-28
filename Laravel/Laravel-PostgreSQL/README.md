# Laravel x PostgreSQL

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend # Laravelプロジェクトのルートディレクトリ
├── infra
│     └── docker
│          ├── nginx
│          │   ├── Dockerfile
│          │   └── default.conf
│          └── php
│              ├── Dockerfile
│              ├── php-fpm.d
│              │   └── zzz-www.conf => unixドメインソケットの設定ファイル
│              └── php.ini
├── Makefile
└── docker-compose.yml
```

## 2. Laravelプロジェクトの新規作成(Makefile)

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
make create-project
```
Webブラウザで http://localhost へアクセスし、Laravelが起動していることを確認する。

## 3. Reactの環境構築(Makefile)

https://qiita.com/morry_48/items/abd620f051fb4f36dcc2 を参考に、下記のコマンドを順に実行する。
```
# パッケージの復元
make npm

# UIパッケージを導入
docker-compose exec app composer require laravel/ui

# Reactを導入
docker-compose exec app php artisan ui react --auth

# npmパッケージのインストール&ビルド
make npm-install
make npm-dev
```

## 参考資料

### Laravel x MySQL

https://qiita.com/ucan-lab/items/5fc1281cd8076c8ac9f4

### Laravel x PostgreSQL

https://qiita.com/sakeafterbeer/items/56cea7e981dacdfc686f<br>
https://lightning-shine.com/post1437/

### Laravel x React

https://qiita.com/morry_48/items/abd620f051fb4f36dcc2<br>
https://qiita.com/mineaki27th/items/ad774a41b7a0a68761bd<br>
https://zenn.dev/hrkmtsmt/scraps/7be4097bc922f1<br>
https://qiita.com/daisu_yamazaki/items/b946594896179abcd203
