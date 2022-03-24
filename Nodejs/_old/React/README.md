# Setup React

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── .gitignore
├── Dockerfile
└── docker-compose.yml
```

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build --no-cache
```

## 3. Reactアプリを作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm frontend /bin/sh -c 'npx create-react-app app --template typescript'
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。

## 4. Dockerfileの編集

`Dockerfile`を下記のように書き換える。
```
FROM node:16.14.2-alpine

ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo

WORKDIR /usr/src/app
```

## 5. Dockerイメージの再ビルド & Dockerコンテナの起動

下記コマンドを実行し、Dockerfileの変更をイメージに反映させ、コンテナを起動する。
```
docker-compose up -d --build
```
Webブラウザで http://localhost:3000 へアクセスし、Reactが起動していることを確認する。

## 5. その他

## 参考資料

https://qiita.com/kashimuuuuu/items/b5f35057dfe1980d053a<br>
https://zenn.dev/daisukesasaki/articles/9620f7fd0ca348<br>
