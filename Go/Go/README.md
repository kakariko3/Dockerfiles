# Setup Go

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── docker-compose.yml
├── Dockerfile
├── Dockerfile.prod
└── main.go
```

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build --no-cache
```

## 3. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```

## 4. その他

## VSCodeの自動補完機能を有効化

VSCodeのサイドバーのDockerロゴをクリックすると、`CONTAINERS`に起動しているコンテナの一覧が表示される。<br>
appコンテナを右クリックし、`Attach Visual Studio Code`を選択して新規ウィンドウを開き、VSCode拡張機能(golang.go)をインストールする。<br>
また、コマンドパレットから`Go: Install/Update Tools`を選択し、拡張機能が依存するGoパッケージをインストールする。<br>

参考:<br>
https://zenn.dev/nkys39/articles/setup-docker-commandless<br>
https://zenn.dev/tomi/articles/2020-10-22-go-docker<br>

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

## Goのプログラムを実行

`/go/src/app`に移動し、下記コマンドを実行する。
```
go run main.go
```

また、コンテナ起動時に`main.go`を実行する場合、`docker-compose.yml`のcommandを以下のように書き換える。
```
command: /bin/sh -c 'go run main.go'
```

## 参考資料

https://zenn.dev/tomi/articles/2020-10-14-go-docker<br>
https://zenn.dev/tatsurom/articles/golang-docker-environment<br>
