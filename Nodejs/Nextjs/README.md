# Next.js Setup

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
- Dockerfile
- docker-compose.yml

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build --no-cache
```

## 3. Next.jsのインストール + アプリ作成

下記コマンドを実行し、Next.jsアプリを作成する。
```
docker-compose run --rm front npx create-next-app next_app --typescript
```
`--rm` : 停止後コンテナを削除<br>
`--use-npm` : パッケージマネージャーにnpmを指定<br>
`--example xxx` : テンプレートを指定

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザで http://localhost:3000 へアクセスし、Next.jsが起動していることを確認する。

## 5. その他

## 参考資料

https://zenn.dev/taku1115/articles/6c9fa97ab37e38
