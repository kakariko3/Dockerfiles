# Setup PHP x PostgreSQL x Apache

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── db
│   └── init
│       └── 0001_init.sql
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
https://zenn.dev/zmb/articles/ffddef3ecd406c<br>
https://qiita.com/sakeafterbeer/items/56cea7e981dacdfc686f<br>
https://zenn.dev/udon_tabee/articles/b6adef86fea894<br>
https://crudzoo.com/blog/docker-postgres<br>
