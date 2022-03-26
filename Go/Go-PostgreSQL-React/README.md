# Setup Go-PostgreSQL-React

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend
│   ├── Dockerfile
│   ├── Dockerfile.prod
│   └── main.go
├── frontend
│   └── Dockerfile
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
`--template typescript` : TypeScriptを利用するためのテンプレートを指定。<br>

下記コマンドを実行し、Reactアプリを`app`ディレクトリから`frontend`ディレクトリ直下に移動する。
```
mv frontend/app/{*,.*} frontend
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル

下記コマンドを実行し、空になったappディレクトリを削除する。
```
rmdir frontend/app
```

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザで http://localhost:3000 へアクセスし、Reactが起動していることを確認する。

## 5. その他

## VSCodeの自動補完機能を有効化

https://zenn.dev/tomi/articles/2020-10-22-go-docker<br>
上記のURLを参考に、コンテナ(backend)にリモートアクセスする新規ウィンドウを開き、VSCode拡張機能(golang.go)をインストールする。<br>
また、コマンドパレットから`Go: Install/Update Tools`を選択し、拡張機能が依存するGoパッケージをインストールする。

## Go Modulesでパッケージをモジュールとして管理

`/go/src/app`に移動し、下記コマンドを実行する。
```
go mod init app
```
Go Modulesの初期化が行われ、`go.mod`が作成される。<br>

パッケージをインストールする場合、importに必要なパッケージを記述し、下記コマンドを実行する。
```
go mod tidy
```
`go.mod`及び`go.sum`にパッケージの情報が反映される。<br>

参考:<br>
https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode<br>
https://nishinatoshiharu.com/go-modules-overview/<br>

## ESLint & Prettier の設定

下記URLを参考に、ESLintとPrettierのインストールと設定を行う。

参考:<br>
https://zenn.dev/jpn_asane/articles/d7f44682b74fdc<br>
https://yumegori.com/vscode_react_typescript_eslint_prettier<br>

## 参考資料

Go:
https://zenn.dev/tomi/articles/2020-10-14-go-docker<br>
https://zenn.dev/tatsurom/articles/golang-docker-environment<br>

React:
https://qiita.com/kashimuuuuu/items/b5f35057dfe1980d053a<br>
https://zenn.dev/daisukesasaki/articles/9620f7fd0ca348<br>
