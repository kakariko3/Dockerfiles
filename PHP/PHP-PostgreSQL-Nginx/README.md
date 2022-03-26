# Setup PHP x PostgreSQL x Nginx

## 1. 作業ディレクトリの作成、各種ファイルの準備

任意の名前のディレクトリを作成し、そのディレクトリ直下に下記のとおりファイルを配置する。
```
.
├── db
│   └── initdb
│       └── init.sql
├── html
│   └── index.php
├── web
│   ├── Dockerfile
│   ├── httpd.conf
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

## 7. その他

## 参考資料

https://qiita.com/ucan-lab/items/5fc1281cd8076c8ac9f4<br>
https://qiita.com/kouki_o9/items/113002580945e94b1456<br>
https://www.ritolab.com/entry/220<br>
