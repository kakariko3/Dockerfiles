# React Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build
```

## 3. Reactのインストール + アプリ作成

下記コマンドを実行し、Reactアプリを作成する。
```
docker-compose run --rm front npx create-react-app react_app --template typescript
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up
```
Webブラウザで http://localhost:8000 へアクセスし、Reactが起動していることを確認する。

## 5. その他

## エラーが起きた場合の対処

`git clone`をしてアプリを起動する場合、下記の対処を行う。
```
docker-compose run --rm front /bin/sh -c "cd react_app && yarn install"
```
`/アプリ名/react_app/node_modules/react-scripts/scripts/utils/verifyTypeScriptSetup.js`内のファイルを下記サイトを参考に書き換える。<br>
https://qiita.com/ke1t0/items/54fb5886439775f20d93

## 参考資料

https://blog.web.nifty.com/engineer/2714<br>
https://qiita.com/2754github/items/413bdaaa90834e219dc8
