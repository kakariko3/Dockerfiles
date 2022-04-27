# Setup Next.js

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

## 3. Next.jsのインストール + プロジェクトの作成

下記コマンドを実行し、Next.jsプロジェクトを作成する。
```
docker-compose run --rm frontend npx create-next-app test --typescript
```
`--rm` : 停止後コンテナを削除<br>
`typescript` : TypeScriptを利用するためのテンプレートを指定。<br>
`--use-npm` : パッケージマネージャーにnpmを指定<br>
`--example xxx` : テンプレートを指定<br>

下記コマンドを実行し、Next.jsプロジェクトを`test`ディレクトリからルートディレクトリ直下に移動する。
```
mv test/{*,.*} .
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル<br>

下記コマンドを実行し、空になったtestディレクトリを削除する。
```
rmdir test
```

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザで http://localhost:3000 へアクセスし、Next.jsが起動していることを確認する。

## 5. その他

## ESLint & Prettier の設定

Prettierをインストールする。
```
yarn add --dev prettier eslint-config-prettier
```

下記URLを参考に、ESLintとPrettierの設定ファイルを構築する。

参考:<br>
https://zenn.dev/hungry_goat/articles/b7ea123eeaaa44<br>
https://qiita.com/kewpie134134/items/0298e5b7a88a06804cd8<br>
https://fwywd.com/tech/next-eslint-prettier<br>

## 参考資料

https://zenn.dev/tasuya/articles/da033574b85e6d<br>
