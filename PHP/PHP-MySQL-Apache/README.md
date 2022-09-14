# Setup PHP x MySQL x Apache

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── db
│   ├── init
│   │   └── 0001_init.sql
│   └── my.cnf
├── html
│   └── index.php
├── web
│   ├── Dockerfile
│   ├── 000-default.conf
│   └── php.ini
├── .gitignore
└── docker-compose.yml
```

## 2. Dockerコンテナの起動

下記コマンドを実行し、コンテナを起動する。
```
docker-compose up -d
```
Webブラウザを起動して http://localhost:8080/ へアクセスし、アプリが起動することを確認する。

## 3. その他

## 参考資料

https://qiita.com/ucan-lab/items/38cd04cee1f3f9e024b9<br>
https://qiita.com/me-654393/items/1ed212cb33eafe734835<br>
https://zenn.dev/ikuraikura/articles/d166724f2c123d1db312<br>
https://zenn.dev/aono/articles/0856cf4d4e3e25<br>
https://tech.fusic.co.jp/posts/2021-08-12-php8-lamp-on-docker/<br>
https://qiita.com/maip0902/items/ae77d2516ce3f3bcef2d<br>
