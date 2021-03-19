# Anaconda Setup

## 1. Dockerイメージの作成
```
$ docker build -t anaconda .
```

## 2. Dockerコンテナの起動
```
$ docker run -itd -v <作業ディレクトリ>/work --name <コンテナ名> anaconda
```

## 3. その他
### bashを起動
---
```
$ docker exec -it <コンテナ名> /bin/bash
```
