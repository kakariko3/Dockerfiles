# Setup Vue.js

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前の作業ディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── docker
│   └── frontend
│       └── Dockerfile
├── .gitignore
└── docker-compose.yml
```

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。
```
docker-compose build --no-cache
```

## 3. Vueプロジェクトの作成

下記コマンドを実行し、Vueプロジェクトを作成する。
```
docker-compose run --rm frontend npx vue create .
```
下記URLを参考に、Vueプロジェクトの各オプションを選択する。<br>
https://qiita.com/shizen-shin/items/a2521440b7dbcab65f9d

## 4. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザで http://localhost:8080 へアクセスし、Vue.jsが起動していることを確認する。

## 5. その他

## 参考資料

https://zenn.dev/chida/articles/8f16e42364398c<br>
https://zenn.dev/tasuya/articles/ad5d71c46db516
