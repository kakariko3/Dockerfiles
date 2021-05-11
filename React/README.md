# React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
$ docker-compose build
```

## 3. Reactのインストール + アプリ作成

下記コマンドを実行し、Reactアプリを作成する。
```
$ docker-compose run --rm node npx create-react-app react_app --template typescript
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
$ docker-compose up
```
Webブラウザで http://localhost:8000 へアクセスし、Reactが起動していることを確認する。

## 5. その他

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
https://blog.web.nifty.com/engineer/2714<br>
https://qiita.com/2754github/items/413bdaaa90834e219dc8
