# Setup PHP x PostgreSQL x Vite

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── backend
│   ├── php-fpm.d
│   │   └── zzz-www.conf
│   ├── src
│   │   └── index.php
│   ├── Dockerfile
│   └── php.ini
├── db
│   └── initdb
│       └── init.sql
├── frontend
│   └── Dockerfile
├── nginx
│   ├── default.conf
│   └── Dockerfile
├── .gitignore
└── docker-compose.yml
```

## 2. Dockerイメージのビルド

ターミナルを開いて作業ディレクトリに移動し、下記コマンドを実行する。<br>
```
docker-compose build --no-cache
```

## 3. Viteプロジェクトを作成

下記コマンドを実行し、Viteプロジェクトを作成する。
```
docker-compose run --rm frontend /bin/sh -c 'yarn create vite test'
```
`--rm` : 停止後コンテナを削除<br>

下記コマンドを実行し、Viteプロジェクトを`frontend/test`ディレクトリから`frontend`ディレクトリ直下に移動する。
```
mv frontend/test/{*,.*} frontend
```
`*` : 全てのディレクトリ・ファイル<br>
`.*` : 全てのドットディレクトリ・ドットファイル

下記コマンドを実行し、空になったtestディレクトリを削除する。
```
rmdir frontend/test
```

## 4. Viteの設定ファイルの作成

appディレクトリ直下に'vite.config.js'を作成し、下記のように記述する。
```
import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    host: '0.0.0.0',
  },
});
```

## 5. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザで http://localhost:3000 へアクセスし、Viteが起動していることを確認する。

## 6. その他

## ESLint & Prettier の設定

下記URLを参考に、ESLintとPrettierのインストールと設定を行う。

参考:<br>
https://zenn.dev/jpn_asane/articles/d7f44682b74fdc<br>
https://yumegori.com/vscode_react_typescript_eslint_prettier<br>

## 参考資料

https://qiita.com/ucan-lab/items/5fc1281cd8076c8ac9f4<br>
https://qiita.com/kouki_o9/items/113002580945e94b1456<br>
https://www.ritolab.com/entry/220<br>
