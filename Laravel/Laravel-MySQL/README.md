# Laravel x MySQL Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend # Laravelプロジェクトのルートディレクトリ
├── infra
│     └── docker
│          ├── mysql
│          │   ├── Dockerfile
│          │   └── my.cnf
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

## 2. Laravelプロジェクトの作成

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
make create-project
```

## 3. その他

## 参考資料

https://qiita.com/ucan-lab/items/5fc1281cd8076c8ac9f4
