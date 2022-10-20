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
```sh
docker-compose build --no-cache
```

## 3. Reactプロジェクトの作成

下記コマンドを実行し、Reactプロジェクトを作成する。
```sh
# using npm
docker-compose run --rm frontend /bin/sh -c 'npx create-react-app test --template typescript'

# using yarn
docker-compose run --rm frontend /bin/sh -c 'yarn create react-app test --template typescript'
```
`--rm` : 停止後コンテナを削除<br>
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。<br>

下記コマンドを実行し、Reactプロジェクトを`app/test`ディレクトリから`app`ディレクトリ直下に移動する。
```sh
mv app/test/{*,.*} app
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル

下記コマンドを実行し、空になったtestディレクトリを削除する。
```sh
rmdir app/test
```

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```sh
docker-compose up -d
```
Webブラウザで http://localhost:3000 へアクセスし、Reactが起動していることを確認する。

## 5. その他

## ESLint & Prettier の設定

下記URLを参考に、ESLintとPrettierのインストールと設定を行う。

参考:<br>
https://zenn.dev/jpn_asane/articles/d7f44682b74fdc<br>
https://yumegori.com/vscode_react_typescript_eslint_prettier<br>
https://zenn.dev/sikkim/articles/93bf99d8588e68<br>

## 参考資料

https://qiita.com/kashimuuuuu/items/b5f35057dfe1980d053a<br>
https://zenn.dev/daisukesasaki/articles/9620f7fd0ca348<br>
