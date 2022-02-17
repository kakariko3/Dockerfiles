# Setup Go-Learning

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── app
│   └── main.go
├── docker-compose.yml
└── Dockerfile
```

## 2. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```

下記コマンドを実行し、appコンテナに入る。
```
docker-compose exec app ash
```

## 3. その他

## VSCodeの自動補完機能を有効化

https://zenn.dev/tomi/articles/2020-10-22-go-docker<br>
上記のURLを参考に、コンテナ(app)にリモートアクセスする新規ウィンドウを開き、VSCode拡張機能(golang.go)をインストールする。<br>
また、コマンドパレットから`Go: Install/Update Tools`を選択し、拡張機能が依存するGoパッケージをインストールする。

## Go Modulesでパッケージをモジュールとして管理

appディレクトリに移動し、下記コマンドを実行する。
```
go mod init app
```
カレントディレクトリ直下に`go.mod`が作成される。<br>

パッケージをインストールする場合、importに必要なパッケージを記述し、下記コマンドを実行する。
```
go mod tidy
```
`go.mod`及び`go.sum`にパッケージの情報が反映される。<br>

参考: 
https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode<br>
https://nishinatoshiharu.com/go-modules-overview/<br>

## 参考資料

https://zenn.dev/tomi/articles/2020-10-14-go-docker<br>
